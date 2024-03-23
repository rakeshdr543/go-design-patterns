package observer

type IObserver interface {
	Update(string)
	GetId() string
}
