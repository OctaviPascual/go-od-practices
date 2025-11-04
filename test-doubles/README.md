# Test doubles

A test double is an object that stands in for a real implementation during testing.

There are multiple types of test doubles:

- Dummy: object that is passed around but never actually used. It's usually not needed in Go because if you know that an interface won't be used you can just use `nil`
- Stub: object that has a fixed set of canned responses
- Fake: working implementation, same behaviour what it replaces but usually with some shortcuts
- Spy: stub that also records information based on how it was called
- Mock: object with predefined expectations that fail the test if they’re not met

## Resources

- https://ieftimov.com/posts/testing-in-go-test-doubles-by-example/
- https://martinfowler.com/bliki/TestDouble.html
- https://quii.gitbook.io/learn-go-with-tests/testing-fundamentals/working-without-mocks#a-primer-on-test-doubles
