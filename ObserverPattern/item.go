package ObserverPattern

import "fmt"

type Item struct {
	observerList []IObserver
	itemName     string
	inStock      bool
}

func NewItem(name string) *Item {
	return &Item{
		itemName: name,
	}
}

func (i *Item) UpdateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.itemName)
	i.inStock = true
	i.notifyAll()
}

func (i *Item) Register(o IObserver) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) Deregister(o IObserver) {
	i.observerList = removeFomObserverList(i.observerList, o)
}

func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.SendNotificationToCustomer(i.itemName)
	}
}

func removeFomObserverList(observerList []IObserver, o IObserver) []IObserver {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observer.GetID() == o.GetID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return nil
}
