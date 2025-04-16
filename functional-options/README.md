# Functional options

This shows how to use functional options, which is a creational design pattern that eases instantiation of objects.


## Simple

- Operate directly on the object

## Advanced

- Use a non-exported intermediate struct to hold the options
- Allow to return an error when applying an option
- Allow to define default values for the options
- Use WithServerOptions to aggregate multiple options to not have to use append

## Opaque interface

- Use opaque interface (interface that exports no methods) to hide serverOptions internal struct

## Resources

- https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
- https://konradreiche.com/blog/how-test-packages-help-to-avoid-interfaces-in-go/
- https://www.youtube.com/watch?v=5uM6z7RnReE
