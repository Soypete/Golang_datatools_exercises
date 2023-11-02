package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"expvar"

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

var (
	countLogin               expvar.Int
	countLogout              expvar.Int
	countPurchase            expvar.Int
	countReturn              expvar.Int
	countAddToCart           expvar.Int
	countRemoveFromCart      expvar.Int
	countViewItem            expvar.Int
	countViewCart            expvar.Int
	countViewPurchaseHistory expvar.Int

	url   string
	topic string
)

func main() {
	flag.StringVar(&url, "url", nats.DefaultURL, "The nats server URLs (separated by comma)")
	flag.StringVar(&topic, "topic", "hello.world", "The nats topic")
	flag.Parse()

	// setup expvar server
	expvar.Publish("countLogin", &countLogin)
	expvar.Publish("countLogout", &countLogout)
	expvar.Publish("countPurchase", &countPurchase)
	expvar.Publish("countReturn", &countReturn)
	expvar.Publish("countAddToCart", &countAddToCart)
	expvar.Publish("countRemoveFromCart", &countRemoveFromCart)
	expvar.Publish("countViewItem", &countViewItem)
	expvar.Publish("countViewCart", &countViewCart)
	expvar.Publish("countViewPurchaseHistory", &countViewPurchaseHistory)

	// setup NATS
	fmt.Println("Connecting to NATS...")
	nc, err := nats.Connect(url)
	if err != nil {
		panic(err)
	}
	time.Sleep(10 * time.Second)

	nc.Subscribe(topic, handleMessage)
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
	switch msg.Body.Action {
	case "login":
		countLogin.Add(1)
		// handle login
	case "logout":
		countLogout.Add(1)
		// handle logout
	case "purchase":
		countPurchase.Add(1)
		// handle purchase
	case "return":
		countReturn.Add(1)
		// handle return
	case "add to cart":
		countAddToCart.Add(1)
		// handle add to cart
	case "remove from cart":
		countRemoveFromCart.Add(1)
		// handle remove from cart
	case "view item":
		countViewItem.Add(1)
		// handle view item
	case "view cart":
		countViewCart.Add(1)
		// handle view cart
	case "view purchase history":
		countViewPurchaseHistory.Add(1)
		// handle view purchase history
	default:
		fmt.Println("Unknown message type")
	}
}
