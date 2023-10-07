# Data Tools for Golang Apps

In Go, datastores are first-class citizens; through the Go standard library, you can leverage database tooling to connect to your datastore using a variety of open source ODBC drivers. You’ll learn about connecting to a relational database, a NoSQL database, or an I/O Filestore and how to transport data between services and manage data contracts. And you’ll dive into testing practices around data tools in Go and learn how to leverage unit and integration tests with a datastore and data contracts.

![wakatime](https://wakatime.com/badge/user/953eeb5a-d347-44af-9d8b-a5b8a918cecf/project/824bacd6-67c2-479a-8e68-c5eee208e5b4.svg)](https://wakatime.com/badge/user/953eeb5a-d347-44af-9d8b-a5b8a918cecf/project/824bacd6-67c2-479a-8e68-c5eee208e5b4)

---

## Database exervices

This section of the course contains examples and exercises for effectively leveraging a Go's database toolset. Databases are a common tool for statemanagement in software engineering, but a lot of the patterns used in interacting with a relational database can be applied to interacting with any data storage system.These exercises are designed to be completed quickly and to set a pattern that can be applied to any software project. They can be completed by following along with the instruction or independently.

### Exercise 1

In your server project, add your preferred database driver and connect to the database in the main function (if you missed day one's exercises or have them in a different location using the [ex-1-connection/main.go](ex-1-connection/main.go)). After you have connected and verified your connection, explore the database. Make sure to query the database's users table and handle the error. Try running `SELECT`, `INSERT`, and `UPDATE` statements

#### Follow up questions:

* what kind of package organization would make sense for organizing your database logic?
* what data base driver did you pick?
* did the data persist?

_NOTE_: If you are not using postgres or are completeling this indepently. You can run many datbases locally using docker. Below is an example of running postgres locally in a docker container.

```
docker pull postgres
docker run -e POSTGRES_PASSWORD=my_password -e POSTGRES_USERNAME=user1 -p 5431:5432 postgres
```

After you get docker running in your local environment setup your database. You will need to `CREATE` your tables and `INSERT` data into the table. You can do this in your Go app or via a sql script editor. [psql](https://www.postgresql.org/docs/current/app-psql.html) is postgres's command line tool.

An example of a go app that connect to a local postgres instance is in [database/ex-1-connection/solution](database/ex-1-connection/solution/postgres.go).

### Live Demo

Using tools to manage a database. Tools like goose and sqlc can be used to easliy abstract database management into your software stack. The following page is the resulting [code from the live demo](database/demo/main.go). The tools in this demo are [sqlc](https://sqlc.dev/) and [goose](https://github.com/pressly/goose). These are very popular in the community.

[Sqlc demo](https://youtu.be/X5VGxx4aQAU)
[goose demo](https://youtu.be/3TnEeRttvyo)

### Exercise 2 Mocking a Database for tests

Interfaces are used to abstract your data storage layer api from the storage technology itself. This will allow you build business technology that is agnostic of your choice of data store. In this exercise, create a `type datastore struct` and a `type dataConnector interface` to mock the database functions into these [api tests.](). The goal is to imitate db interactions without connecting to the db. You will be adding your abstraction logic to [this `database` package](). There is no `func main()` because the execution is done completely through tests

[Here](https://github.com/Soypete/golang-cli-game/blob/24dc57852dee27bb17120555d3d390bd17a78d13/server/api_test.go#L14) are some working tests that use `passBD{}` and `failDB{}` to mock database functionality in an api test.

## Streaming Data Exercises

### Exercise 3 streaming client

### gRPC demo

### Exercise 4 grpc data contract tests

---

## Additional Resources:
