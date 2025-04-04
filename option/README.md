# Option

This is an example on how to use functional options.

## Simple

https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis

## Advanced

https://konradreiche.com/blog/how-test-packages-help-to-avoid-interfaces-in-go/
https://www.youtube.com/watch?v=5uM6z7RnReE

We create a separate struct to hold the options, to not pollute the original struct.
We allow to return an error when applying an option.
When creating Server we create an internal options to be able to specify default options.
WithServerOptions allows to aggregate multiple options into a single one to not have to use append on them.
