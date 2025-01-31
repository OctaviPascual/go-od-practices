# Embedding

This is an example that shows how to take advantage of embedding.

Instead of dropping the final approach, it's worth to show each of the alternatives, from most basic to more advanced to see the improvements that each approach brings to the table. That way it's more evident why we do things that way and not another.

## Intro

We want to wrap a `Cacher` interface provided by an external lib to have a verbose cache, which logs a message every time a key is added and tracks the number of add operations that occurred. We also want this verbose cache to implement the `Cacher` interface to pass it to methods accepting a `Cacher` type.

```
Adding (key="1",value=1) to cache.
Adding (key="2",value=2) to cache.
Total number of adds: 2
```

## Struct

The first approach consists of a struct to represent the cache.

Cons:
- we cannot change the underlying cache from LRU to LFU
- we have to redefine all methods of the `Cacher` interface

## Struct embedding

We can use embedding to improve the situation. Embedding is a mechanism that promotes a field of the struct when it has not an explicit name.

Pros:
- we have to redefine only the methods of `Cacher` that we want to override

Cons:
- we cannot change the underlying cache from LRU to LFU

## Interface

We can use an interface.

Pros:
- we can change the underlying cache from LRU to LFU

Cons:
- we have to redefine all methods of the `Cacher` interface

## Interface embedding

We can use embedding with an interface.

Pros:
- we can change the underlying cache from LRU to LFU
- we have to redefine only the methods of `Cacher` that we want to override

## Resources

This was completely inspired from those 3 posts:
- https://eli.thegreenplace.net/2020/embedding-in-go-part-1-structs-in-structs/
- https://eli.thegreenplace.net/2020/embedding-in-go-part-2-interfaces-in-interfaces/
- https://eli.thegreenplace.net/2020/embedding-in-go-part-3-interfaces-in-structs
- https://fabianlindfors.se/blog/decorators-in-go-using-embedded-structs/