# Generics

These examples walk you through implementations of generics in GO.
To read full design draft : [Design Draft](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md)

## Very High Level Overview :

This section explains the changes suggested by the design draft very briefly. This section is intended for people who are already familiar with how generics would work in a language like Go. These concepts will be explained in detail in the following sections.

* Functions can have an additional type parameter list that uses square brackets but otherwise looks like an ordinary parameter list: func F[T any](p T) { ... }.
* These type parameters can be used by the regular parameters and in the function body.
* Types can also have a type parameter list: type M[T any] []T.
* Each type parameter has a type constraint, just as each ordinary parameter has a type: func F[T Constraint](p T) { ... }.
* Type constraints are interface types.
* The new predeclared name any is a type constraint that permits any type.
* Interface types used as type constraints can have a list of predeclared types; only type arguments that match one of those types satisfy the constraint.
* Generic functions may only use operations permitted by the type constraint.
* Using a generic function or type requires passing type arguments.
* Type inference permits omitting the type arguments of a function call in common cases