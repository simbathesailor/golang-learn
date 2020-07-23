package cfg

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type ChannelNameSpace string

type SubscribedChannel interface {
	ChannelName() ChannelNameSpace
	Client() *redis.PubSub
	Unsubscribe() error
	PollMessage(frequency time.Duration, onMessage func(*redis.Message))
}

type subscribedChannel struct {
	name    ChannelNameSpace
	channel *redis.PubSub
}

func (c *subscribedChannel) ChannelName() ChannelNameSpace {
	return c.name
}

func (c *subscribedChannel) Client() *redis.PubSub {
	return c.channel
}

func (c *subscribedChannel) Unsubscribe() error {
	return c.channel.Unsubscribe(string(c.name))
}

func (c *subscribedChannel) PollMessage(frequency time.Duration, onMessage func(*redis.Message)) {
	go func() {
		ticker := time.NewTicker(frequency)
		for {
			select {
			case <-ticker.C:
				if message, _ := c.channel.ReceiveMessage(); message != nil {
					onMessage(message)
				}
			}
		}
	}()
}

type Cache interface {
	Client() *redis.Client
	Set(key string, value interface{})
	SetWithExpiration(key string, value interface{}, expiration time.Duration)
	Get(key string, value interface{}) error
	Subscribe(channel ChannelNameSpace) SubscribedChannel
	Publish(channel ChannelNameSpace, data interface{}) error
	Close() error
}

type cache struct {
	client             *redis.Client
	subscribedChannels []*subscribedChannel
}

func (c *cache) Client() *redis.Client {
	return c.client
}

func (c *cache) Set(key string, value interface{}) {
	byteArr, _ := json.Marshal(value)
	c.client.Set(key, byteArr, 0)
}

func (c *cache) SetWithExpiration(key string, value interface{}, expiration time.Duration) {
	c.client.Set(key, value, expiration)
}

func (c *cache) Get(key string, value interface{}) error {
	bytes, err := c.client.Get(key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, value)
}

func (c *cache) Subscribe(channelName ChannelNameSpace) SubscribedChannel {
	subChannel := c.client.Subscribe(string(channelName))
	s := &subscribedChannel{
		name:    channelName,
		channel: subChannel,
	}
	c.subscribedChannels = append(c.subscribedChannels, s)
	return s
}

func (c *cache) Publish(channelName ChannelNameSpace, data interface{}) error {
	byteArr, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.client.Publish(string(channelName), byteArr).Err()
}

func (c *cache) Close() error {
	for i := range c.subscribedChannels {
		if err := c.subscribedChannels[i].channel.Close(); err != nil {
			return fmt.Errorf("failed to close subscribed channel : %w", err)
		}
	}
	if err := c.Close(); err != nil {
		return fmt.Errorf("failed to close cache client : %w", err)
	}
	return nil
}

// returns newly instantiated redis client
func initializeCache(config *AppConfig) *cache {
	client := redis.NewClient(&redis.Options{
		Addr: config.Cache.Host,
	})
	return &cache{
		client: client,
	}
}
