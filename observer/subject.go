package observer

type ISubject interface {
	RegisterObserver(o IObserver)
	RemoveObserver(o IObserver)
	NotifyObservers()
}
