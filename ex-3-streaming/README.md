# Exercise 3

In this exercise you will be agregating live event to count the number or user actions that occur. You are provided a database to reference and answer the follow-up question.

Using provided [nats subscriber]() subscriber handle messages that come in. Handle the messages by parsing them into a struct. Using the data in the struct, collect stats about the total number of actions that take place for each action type. Below are a set of questions. Use the parsed data to answer the questions. Run your consumer program for 90 seconds to collect the stats and them print the answers to standard out.

```
type Message struct {
	ID        int
	Type      string
	Body      UserAction
	Timestamp time.Time
}
```

The following action types will be used.

* login
* logout
* purchase
* return
* add to cart
* remove from cart
* view item
* view cart,
* view purchase history

Print the answers to the following questions to your standard out.

1.  How many purchases did \_\_ make?
2.  how many "add item to cart" messages failed?
3.  Of the total message processed how many were of \_\_\_ type?

Solution [here](solution/main.go)

_Note_: if you are following this on your own use docker to create a local nats server. That can be done with the following command:

```
docker run --name nats -p 4222:4222 -p 8222:8222 nats --http_port 8222
```

Make sure to change your connection to the following IP

```
	nc, err := nats.Connect("http://localhost:4222")
	if err != nil {
		panic(err)
	}
```
