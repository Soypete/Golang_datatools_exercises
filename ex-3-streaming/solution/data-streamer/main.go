package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
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

var UserActions = []string{"login", "logout", "purchase", "return", "add to cart", "remove from cart", "view item", "view cart", "view purchase history"}

// MessageCreate creates a message with a random ID, a type, a body, and a timestamp.
// It returns the message as a serialized byte array and an error if there is one.
func MessageCreate(messageType string, messageBody UserAction) ([]byte, error) {
	message := Message{
		ID:        rand.Intn(100000),
		Type:      messageType,
		Body:      messageBody,
		Timestamp: time.Now(),
	}
	return json.Marshal(message)
}

// chooseMessageType chooses a message type at random.
func chooseMessageType() string {
	return UserActions[rand.Intn(len(UserActions))]
}

func (u *UserAction) getErrorMessage() {
	if u.Succeded {
		u.Error = ""
	}
	switch u.Action {
	case "login":
		u.Error = "login failed"
	case "logout":
		u.Error = "logout failed"
	case "purchase":
		u.Error = "purchase failed"
	case "return":
		u.Error = "return failed"
	case "add to cart":
		u.Error = "add to cart failed"
	case "remove from cart":
		u.Error = "remove from cart failed"
	case "view item":
		u.Error = "view item failed"
	case "view cart":
		u.Error = "view cart failed"
	case "view purchase history":
		u.Error = "view purchase history failed"
	default:
		u.Error = "unknown action"
	}
}

// messagePublish publishes a message to a NATS server.
// It returns an error if there is one.
func messagePublish(nc *nats.Conn) error {

	// populate message body
	userAction := UserAction{
		UserID:    rand.Intn(20), // there are 20 users in the database
		Action:    chooseMessageType(),
		Succeded:  rand.Intn(2) == 0,
		Timestamp: time.Now(),
	}
	userAction.getErrorMessage()

	// create message
	message, err := MessageCreate("userAction", userAction)
	if err != nil {
		return err
	}

	fmt.Println(string(topic))
	// publish message
	nc.Publish(topic, message)
	log.Printf("Published message: %s\n", message)
	return nil
}

var (
	url   string
	topic string
)

func main() {
	flag.StringVar(&url, "url", nats.DefaultURL, "The nats server URLs (separated by comma)")
	flag.StringVar(&topic, "topic", "hello.world", "The nats topic")
	flag.Parse()

	fmt.Println("url:", url)
	nc, err := nats.Connect(url)
	if err != nil {
		panic(err)
	}

	// symulate users on many devices
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		for {
			messagePublish(nc)
			time.Sleep(1 * time.Second)
		}
	}()
	wg.Wait()

	// system close call to close all connections
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		sig := <-c
		fmt.Printf("Got %s signal. Aborting...\n", sig)
		wg.Done()
		nc.Close()
		os.Exit(1)
	}()
}
