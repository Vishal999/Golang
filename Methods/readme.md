# Methods

#### Methods and Indirection

Method with value/pointer receivers take either a value or a pointer as the receiver when they are called while a function that takes a value argument must take a value of that specific type and same is true for pointer as well.


### Choosing a value or pointer receiver

2 reasons to use a pointer receiver : 
* Method can modify the value that it receiver points to.
* To avoid copying on each method call. This can be more efficient if the receiver is a large struct, for example.

**Note** In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.