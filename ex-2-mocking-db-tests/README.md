# Exercise 2 - writing tests with a Data store

I/O operations are cumbersome and should be avoided in unit tests. When you are building service that rely on database conenctions this makes writing unit tests and maintaining good test coverage cumbersome. Using the provided [server.go](/server.go) and [server_test.go](/server_test.go) add mock functions that that emulate the database layer. You should add mock fuctions to replicate one successful db response and one failed db response.

Solution [here](/solution/server_test.go)
