# Embedding

This is an example that shows how to take advantage of embedding.

I think that instead of just providing the approach that is the most convenient, it's worth to still show the other approaches to see what are their problems and why one is more convenient than the others.

We want to wrap a `Cacher` interface provided by an external lib to have a verbose cache, which logs a message every time a key is added and tracks the number of add operations that occurred.


## Struct

The first approach consists of a struct to represent the cache.

Cons:
- we cannot easily change the underlying cache from LRU to LFU
- we have to redefine all methods of the `Cacher` interface

## Struct embedding

We can use embedding to improve the situation. Embedding is a mechanism that promotes a field of the struct when it has not an explicit name.

Pros:
- we don't have to redefine all methods of the `Cacher` interface

Cons:
- we cannot easily change the underlying cache from LRU to LFU

## Interface

We can use an interface.

Pros:
- we can easily change the underlying cache from LRU to LFU

Cons:
- we have to redefine all methods of the `Cacher` interface

## Interface embedding

We can use embedding with an interface.

Pros:
- we can easily change the underlying cache from LRU to LFU
- we don't have to redefine all methods of the `Cacher` interface

## Resources

This was completely inspired from those 3 posts:
- https://eli.thegreenplace.net/2020/embedding-in-go-part-1-structs-in-structs/
- https://eli.thegreenplace.net/2020/embedding-in-go-part-2-interfaces-in-interfaces/
- https://eli.thegreenplace.net/2020/embedding-in-go-part-3-interfaces-in-structs