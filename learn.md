Go live server:

https://github.com/codegangsta/gin

Data types in Golang:

Structs:

Mind the names start with capital letter

type Doctor struct {
Number int
CctorName string
Companions []string
}

aDoctor:= Doctor{
Number: 3
ActorName: “David,
Companions []string {
“Liz hurley”,
“Anil”
},

}

//////
Anonymous struct declaration and definition

aDoctor : struct{name string} {name: “Anil”}

Structs are value type

anotherDoctor:= aDoctor

anotherDoctor.name = “sunil”

Result will be
aDoctor.name = “Anil”
anotherDoctor.name = “sunil”

Embedding model with structs:

type Animal struct {
Name string
Origin string

}

type Bird Struct {
Animal
CanFly bool
}

//////// First way
B: = Bird{}

B.Name = “Emu”
B.Origin = “Australia”

B.CanFly= 48

///// Second way:

B:= Bird {
Animal: Animal{Name: “Emu”, “Origin”: “Austraila”},
CanFly: false

}

Tags

type struct {
Name string `required max: “100”`
Origin string `required max: “50”`

}

Import “reflect”

T:= reflect.TypeOf(Animal{})

field, \_ := t.FieldByName(“Name”)

fmt.Println(field.Tag)

Maps:

v := make([string]int)
v = map[string]int{
“A”: 1,
“B”: 2  
}

Access v[“A”]
Update v[“B”] = 8
delete(v, “A”)
If v[A] is not there, then it will give 0
Return order of map is not guaranteed

\_, ok := v[A]
if(!ok) {
fmt.Println(“not there”, ok)
}

len(v) for length for the map

These are passed by reference.

Queries

package main

import "fmt"

type CoffeeMachine struct {
NumberOfCoffeeBeans int
}

func NewCoffeeMachine() \*CoffeeMachine {
return &CoffeeMachine{}
}

func (cm \*CoffeeMachine) SetNumberOfCoffeeBeans(n int) {
cm.NumberOfCoffeeBeans = n
}

func main() {
cm := NewCoffeeMachine()
fmt.Println(cm)
cm.SetNumberOfCoffeeBeans(100)

    fmt.Printf("The coffee machine has %d beans\n", cm.NumberOfCoffeeBeans)

}

If Statement

If pop, ok := simecheck(); ok {
fmt.Println(pop)
Pop variable is only accessible inside it, not outside
}
fmt.Println(pop) // it will be error

Switch Statement

Switch with tag here 2 is the tag which will be checked in switch.
Switch 2 {
case 1:
fmt.Println(“one”)
case 2:
fmt.Println(“two”)
default :
fmt.Println()  
}

// two will get printed out

case 1, 5, 10:  
 fmt.Println(“It will check 1, 5 10”)

/////
switch i:= 2 + r; i {
Case 1,5,10
}

///////
Switch without tag

Without tags, first pass will run only

It has implicit breaks
i:=10
Switch {
Case i <=10:
Dada
saasdad

Case i <=20:
sdadsds
}

Fallthrough
i:=10
Switch {
Case i <=10:
Dada
saasdad
fallthrough
Case i <=20:
sdadsds
}

Both the condition when i < 10 and also i <=20

var i interface {} = 1

Switch i.(type) {
case int:
Sadsds
  
 case float64:
Asdsads
case string:
Sdadsad
default:
  
 fmt.Print

}

You can aslo make use of break keyword to coume out of switch

For loop
i is scoped to only for loop
for i:=0; i<5;i++ {

}

for i, j:=0, 0; i<5;i, j = i + 1, j+ 1 {

}

i:=0
for ; i<5;i++ {

}

i is scoped to main function above in this specific case

For range loop

s:= []int{1,2,3}

for k, v: range s {
fmt.Println(k, v)
}

This way you can also loop map

Ansn also string

s:= “sdadad”

for k, v: range s {
fmt.Println(k, string(v))
}

DE
Defer, Panic and Recover

Defer

It executes the statement after the last statement is excited in wrapped function and then call the defer statement before returning

Defer follow stack order

Defer dad /// 1
Defer sdadc // 2
Defer sddcsd //3

3/2/1

/////

a:= “start”
defer fmt.Println(a) // start
a = “end”

Panic

a, b = 1, 0
ans:= a/b
fmt.Println(ans)

Will throw error

panic(“something bad happened”)

Panic always run after the defer statements

defer func() {
}()

Recover read more about it, but it like catch statement in javascript
Pointer

Pointer arithmetic is not allowed in golang directly
Pointers arithmetic is available in ‘unsafe’ package

package main

import "fmt"

type A struct {
B int
E string
}
func main() {
c:= A{2, "anil"}
d:= &c
fmt.Println(d)
}

// output &{2 anil}
Just a representation telling it is a pointer holding on two this value

Var ms \* mystruct
Ms = new(mystruct)
Ms.foo = 42
fmt.Println(ms.foo)

Var ms * mystruct
Ms = new(mystruct)
(*ms).foo = 42
fmt.Println((\*ms).foo)

Slice and Array

In slices changes the original array
But in case of array we change the copy.

Slices and maps are like object passed around
But arrays are passed by deep immutable new structure

// Pending new operator

Functions

Variadic parameters

func sum(msg string,values ...int) {
Result : =0
for \_, v:=range values {
result +=v
}
fmt.Println()
}

Panic means application cannot continue

This is also valid

var f func() = func() {
}

f()

Var divide func(float64 , float64) (float64, error)

divide = func(afloat64 , bfloat64) (float64, error) {

}

Type greeter struct {
greeting string
name string

}

// Copy example
Value reciever
func (g greeter) greet() {
fmt.Println(g)
}

Pointer reciever
func (g greeter) greet() {
fmt.Println(g)
}

Using pointer receiver is good in case , when you need to change the original object.

Interfaces

They describe behaviour not data
Import “fmt”

type Writer interface {
Write([]byte) (int, error)
}

In go we implicitly implement interfaces, there is not keyword like implements

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
n, err = fmt.Println(string(data))
Return n, error
}

In Main method
Var w Writer = ConsoleWriter{}
w.Write([]byte(“Hello Go!”))

Append er to interface names

Anything ending with er is mostly interface

We can compose two or more interface and declare a new interface

Var myObj interface{} = NewBufferWriterCloser()

For empty interface , you need to use type case or type conversion

myObj.(WriterCLoser)

…….

Interface type conversion
Bwc, ok := wc.(\*BufferWriterCOloser)

The object with having implementation based on a same interface can be converted interchangeably

The method receiver also need to be same when doing the conversion

Var myObj interface{} = NewBufferedWWriteCloser()

Types switches

Var i interface{} = 0

Switch i.(type) {
Case int:
asSSD
CASE STRING:
SASADsadasdd
default :
csacadad
}

https://youtu.be/YS4e4q9oBaU?t=19241 for specidal behaviour of interface amd methods

Types methods each one of methods regardless of receiver types

Implement interface with concrete value :
Method set of value, any method set with value as a receiver. And it makes method set is incomplete.

If I am implementing an interface, if i am using a value type, the method that implements the interface have to have the value receivers

If I am implementing an interface, if i am using a pointertype, the method can be there regardless of receiver types

io.Writer, io.Reader, interface{}

GoRoutines:

func main() {
go sayHello()
}

func sayHello() {
fmt.Println(“sdasd”)
}

Green thread

Above code will not print the

func main() {
go sayHello()
time.Sleep(100\* time.Millisecond)
}

func sayHello() {
fmt.Println(“sdasd”)
}

Above is one of the hack

func main() {
go func() {
fmt.Println(“sdasd”)

}()
time.Sleep(100\* time.Millisecond)
}

func main() {
Var msg = “hello”
go func() {
fmt.Println(msg)

}()
msg = “opjo”

time.Sleep(100\* time.Millisecond)
}

Opjo will get printed
////////

Var wg = sync.WaitGroup{}

func main() {
Var msg = “hello”
wg.Add(1)
go func() {
fmt.Println(msg)
wg.Done()

}()
msg = “opjo”
wg.Wait()
}

Mutex

Var m = sync.RWMutex{}

Import “runtime”

runtime.Gomaxprocs(100)

Best practices:
Donot create go routines for library

Know how it will end
Let the consumer of the library set the no of threads

Debugiging

go run -race pathtofile.go

Channels

Var wg = sync.WaitGroup()
Func main() {
ch: make(chan, int)
wg.Add(2)
Go func() {
i:= <- ch
fmt.Println(i)
wg.Done()
}()

Go func() {
i:= 42
Ch <- i
fmt.Println(i)
wg.Done()
}()

}
wg.Wait()

Concept of receiver anf sender goroutine

Send only and receive only channel

Buffered channel

ch: make(chan int, 50)

ch:= make(chan int, 50)

wg.Add(2)

// Receive only channel
go func(ch <- chan int) {
i:= <-ch
fmt.Println(i)
wg.Done()
}(ch)

// Send only channel
go func(ch chan<- int) {
Ch <-42
Ch <-27
fmt.Println(i)
wg.Done()
}(ch)

wg.Wait()

// Make use of for looping strategy

// Receive only channel
go func(ch <- chan int) {
i:= <-ch
for i:= range ch{
fmt.Println
}

wg.Done()
}(ch)

// Send only channel
go func(ch chan<- int) {
Ch <-42
Ch <-27
fmt.Println(i)
wg.Done()
}(ch)

wg.Wait()

In case of channel range the first

for i:= range ch{
fmt.Println
}

I is the value itself

You can close the channel with
close(ch)

You cannot send a message to a closed channel

/// on receiver side

for {
If i, ok := <-ch;ok {
fmt.Println(i)
} else {
break
}
