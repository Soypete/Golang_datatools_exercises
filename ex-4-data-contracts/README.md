# Exercise 4: Using gRPC to validate data contract

In this exercise we will be creating and validating Data contracts. We will be using the [protobuf]() + [grpc]() infrastructure to add a first layer of validation. This is a great way to restrict all payloads into a predictable and structured format.

_Note_: This can also be accomplished using [JSON schema](). gRPC is more popular amoung go developers, but there is no conceptual advatange in using `gRPC` over `JSON schema`.

We are using the same data structures as in the previous streaming examples:

```
type Message struct {
	ID        int
	Type      string
	Body      UserAction
	Timestamp time.Time
}

// action types: login, logout, purchase, return, add to cart, remove from cart, view item, view cart, view purchase history
type UserAction struct {
	UserID    int
	Action    string
	Succeded  bool
	Error     string
	Timestamp time.Time
}
```

To complete this exercise, add the required logic to the functions found in [solution.go](/server.go) to make the test pass. The tests are already implemented in [server_test.go](/server_test.go). The tests will start as all failing. You can check your test output by running

```
$ go test /ex-4-data-contracts -v
```

[Solution](/solution/server.go)

Follow-up Questions:

* What fields are required according to our current data contract?
* Should any additional fields be required?
