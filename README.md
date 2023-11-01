# Golang_datatools_exercises

exercises, activities, and questions for integrating data tools to your go program

## [![wakatime](https://wakatime.com/badge/user/953eeb5a-d347-44af-9d8b-a5b8a918cecf/project/815add1c-01f3-412e-b6cd-730805338e0e.svg)](https://wakatime.com/badge/user/953eeb5a-d347-44af-9d8b-a5b8a918cecf/project/815add1c-01f3-412e-b6cd-730805338e0e)

## pre-requisites

* [Go](https://go.dev/) installed and running
* Working knowledge of go
* git for cloning and using the repo

## New to go?

If you are new to go, work through these exercises first

* [Golang Zero to Hero](https://github.com/Soypete/Golang_tutorial_zero_to_hero)
* [Tour of Go](https://go.dev/tour/welcome/1)
* [Gophercises](https://gophercises.com/)

---

## Exercises

### [Exercise 1](/databases/README.md) - Connect a Database to your Go Service

### [Exercise 2](/databases/README.md) - Databases in tests

I/O operations are cumbersome and should be avoided in unit tests. When you are building service that rely on database conenctions this makes writing unit tests and maintaining good test coverage cumbersome. Using the provided program and functions fix the tests to pass by using a DB mock.

Solution [here](../restful-go/ex-4-tests/framework_test.go)

### [Exercise 3](/streaming/READEME.md) - connect a streaming client

In this solution you will be agregating live event to count the number or user actions that occur. You are provided a database to reference and answer the follow-up question.

Using provided [nats subscriber]() subscriber handle messages that come in. Handle the messages by parsing them into a struct. Using the data in the struct, collect stats about the total number of actions that take place for each action type. Below are a set of questions. Use the parsed data to answer the questions. Run your consumer program for 90 seconds to collect the stats and them print the answers to standard out.

```
type Message struct{

}
```

The following action types will be used.

```

```

Print the answers to the following questions to your standard out.

1.  How many purchases did \_\_ make?
2.  how many "add item to cart" messages failed?
3.  Of the total message processed how many were of \_\_\_ type?

Solution [here]()

_Note_: if you are following this on your own use docker to create a local nats server. That can be done with the following command:

```

```

Make sure to change your connection to the following IP

```
	nc, err := nats.Connect("http://localhost:4222")
	if err != nil {
		panic(err)
	}
```

---

## Live Demo

Using tools to manage a database. Tools like goose and sqlc can be used to easliy abstract database management into your software stack. The following page is the resulting [code from the live demo](database/demo/main.go)

---

## Explore More

* [pgx](https://github.com/jackc/pgx)
* [pg](https://github.com/lib/pq)
* [sqlx](https://github.com/jmoiron/sqlx)
* [sqlc](https://sqlc.dev/)
