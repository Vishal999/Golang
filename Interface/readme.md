# Interface 

#### Interface are implemented implicitly
A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

Implicit Interfaces decouple the definition of an interface from its implementation, which coud then appear in any package without prearrangement

**Note**
##### Interface Values
1. A interface value holds the value of a specific underlying concrete type.
2. Calling a method on an interface value executes the method of same name on its underlying type.
3. If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
4. Go gracefully handles calling a method with nil receiver

```golang
// i interface
// t struct that implements i interface
var i I
var t T 
i = T
// nil receiver calling method M
i.M()
```
5. An interface value that holds a nil concrete value is itself non-nil.

##### NIL Interface Values

A nil interface value holds neither value nor concrete type.
Calling a method on a nil interface is a runtime error because there is no type inside the interface tuple to indicate which concrete method to call.

##### Type Assertions 

It provides access to an interface value's underlying concrete type.

To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and the boolean value that reports whether the assertion succeeded.

```golang
t, ok := i.(T)
```

If i holds a T then t will be underlying value and ok will be true.
If not, ok will be false and t will be zero value of type T and no panic occurs.


##### Type Switch

It is a construct that permits several type assertions in series.

It is like a regular switch statement, but the case in a type switch specify type (not values) and those values are compared against the type of value held by given interface value.

```golang
switch v := i.(type){
    case T:
        // Here v has type T
    case S:
        // Here v has type S
    default:
        // no match, here v has same type as i
}

// Golang doesn't allow i.(type) without switch.
v:= i.(type) // panic - use of .(type) outside type switch
```

##### Stringers

```golang
type Stringer interface{
    String() string
}
```

A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.

##### Error

Go program express error state with error values.

The error type is built in interface similar to fmt.Stringer

```Golang
type Error interface{
    Error() string
}
```
(As with fmt.Stringer, the fmt package looks for the error interface when printing values.)

Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.

```golang
i, err := strconv.Atoi("42")
if err!=nil{
    fmt.Printf("coudn't convert number : %v\n", err)
    return
}
fmt.Println("Converted integer : ", i)
```

##### Reader 

The io package specifies io.Reader interface, which represents read end of a stream of data.

The io.Reader interface has a Read method:

```Golang
func (T)Read(b []byte) (n int, err error)
```

Read populates the given byte slice with data and returns number of bytes populated and an error value. It returns an io.EOF error when the stream ends.
