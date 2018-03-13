package main

import (
	"github.com/ardanlabs/gotraining/topics/go/design/composition/mocking/example1/pubsub"
)

// publisher is an interface to allow this package to mock the pub/sub
type publisher interface {
	Publish(key string, v interface{}) error
	Subscribe(key string) error
}

// mock is a concrete type to help support the mocking of the pub/sub package
type mock struct {}

// Publish implements the publisher interface for the mock
func (m *mock)Publish(key string, v interface{}) error {
	// ADD YOUR MOCK FOR THE PUBLISH CALL.
	return nil
}

// Subscribe implements the publisher interface for the mock.
func (m *mock)Subscribe(key string) error {
	// ADD YOUR MOCK FOR THE SUBSCRIBE CALL.
	return nil
}

func main() {
	// Create a slice of publisher interface values.Assign the address of a
	// pubsub.PubSub value and the address of a mock value.
	pubs := []publisher{
		pubsub.New("localhost"),
		&mock{},
	}


	// Range over the interface value to see how the publisher
	for _, p := range pubs{
		p.Publish("key", "value")
		p.Subscribe("key")
	}
}
