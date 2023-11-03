# Golang_datatools_exercises

exercises, activities, and questions for integrating data tools to your go program

## [![wakatime](https://wakatime.com/badge/user/953eeb5a-d347-44af-9d8b-a5b8a918cecf/project/815add1c-01f3-412e-b6cd-730805338e0e.svg)](https://wakatime.com/badge/user/953eeb5a-d347-44af-9d8b-a5b8a918cecf/project/815add1c-01f3-412e-b6cd-730805338e0e)

## Pre-requisites

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

* [Exercise 1](/ex-1-connection/README.md) - Connect a Data base to your Go Service
* [Exercise 2](/ex-2-mocking-db-test/README.md) - Databases in tests
* [Exercise 3](/ex-3-streaming/READEME.md) - connect a streaming client
* [Exercise 4](/ex-4-data-contracts/README.md) - validate your data contract with grpc

---

## Live Demos

### DB Management

Using tools to manage a database. Tools like goose and sqlc can be used to easliy abstract database management into your software stack. The following page is the resulting [code from the live demo](database-demo/main.go)

[watch video]()

### grpc-gateway

Grpc is a valuable tool for managing data contracts. You can leverage grpc-gateway to still maintain restful best practices and accept json payloads while using protobufs to maintain structured data formats. You can view the demo code in [grpc-gateway-demo](/grpc-gateway-demo/main.go)

---

## Explore More

* [pgx](https://github.com/jackc/pgx)
* [pg](https://github.com/lib/pq)
* [sqlx](https://github.com/jmoiron/sqlx)
* [sqlc](https://sqlc.dev/)
* [gRPC](https://grpc.io/)
* [protobuf](https://protobuf.dev/)
* [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
