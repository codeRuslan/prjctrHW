package observer

import "fmt"

type Customer struct {
	ID string
}

// NewCustomer creates a new instance of Customer.
func NewCustomer(id string) *Customer {
	return &Customer{
		ID: id,
	}
}

// Update is the method required by the Observer interface.
func (c *Customer) Update(message string, userJoined string) {
	// Implement the logic for handling the update event.
	// For example, print a message here.
	fmt.Printf("[EVENT]Customer %s: Update for room: %s that %s joined\n", c.ID, message, userJoined)
}

// GetID is the method required by the Observer interface.
func (c *Customer) GetID() string {
	return c.ID
}
