package ObserverPattern

type Observable interface {
	Register(observer IObserver)
	Deregister(observer IObserver)
	NotifyAll()
}
