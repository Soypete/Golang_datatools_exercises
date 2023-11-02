package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
)

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

func main() {

	// setup NATS
	fmt.Println("Connecting to NATS...")
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		panic(err)
	}
	time.Sleep(10 * time.Second)

	nc.Subscribe("hello.world", handleMessage)
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	go http.ListenAndServe(":8080", nil)

	runtime.Goexit()
}

func handleMessage(m *nats.Msg) {

	var msg Message
	err := json.Unmarshal(m.Data, &msg)
	if err != nil {
		log.Println(err)
		return
	}
	// TODO: aggregate data here and collect stats
}
