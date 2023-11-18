package ObserverPattern

import "fmt"

func NewCustomer(e string, n string) customer {
	return customer{
		email: e,
		name:  n,
	}
}

type customer struct {
	email string
	name  string
}

func (c customer) SendNotificationToCustomer(item string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.name, item)
}

func (c customer) GetID() string {
	return c.name
}
