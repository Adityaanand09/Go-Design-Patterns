package ObserverPattern

type IObserver interface {
	SendNotificationToCustomer(string)
	GetID() string
}
