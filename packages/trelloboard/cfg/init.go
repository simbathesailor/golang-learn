package cfg

import "fmt"

var app App

func CreateApp(region, environment string) App {
	fmt.Println("Start creating app")

	// IntializeDatabase()

	// var err error
	// var dbs []*DataStore
	// var config *AppConfig

	// // initialise cfg -- read from cfg sources
	// if config, err = generateAppConfig(region, environment); err != nil {
	// 	panic(err)
	// }

	// // initiate Logger
	// logger := initLogger(config.Environment)

	// // instantiate repository
	// if dbs, err = initDatabases(logger, config); err != nil {
	// 	panic(err)
	// }

	// // initialize redis
	// cache := initializeCache(config)

	// instantiate http px
	httpClient := initHTTP()
	fmt.Println("Here is your http client", httpClient)
	// instantiate server
	// server := createNewServer()

	// app = &appInstance{
	// 	dBs:    dbs,
	// 	logger: logger,
	// 	config: config,
	// 	http:   httpClient,
	// 	server: server,
	// 	cache:  cache,
	// }

	return app
}

// func GetApp() App {
// 	return app
// }
