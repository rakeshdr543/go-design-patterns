package observer

import "fmt"

type Item struct {
	customers []IObserver
	name      string
	inStock   bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) UpdateAvailability() {
	fmt.Println("Item status changed to Available")
	i.inStock = true
	i.NotifyObservers()
}

func (i *Item) Register(o IObserver) {
	i.customers = append(i.customers, o)
}

func (i *Item) Deregister(o IObserver) {
	i.customers = removeFromSlice(i.customers, o)
}

func (i *Item) NotifyObservers() {
	for _, observer := range i.customers {
		observer.Update(i.name)
	}
}

func removeFromSlice(customers []IObserver, observerToRemove IObserver) []IObserver {
	customersLength := len(customers)
	for i, observer := range customers {
		if observer.GetId() == observerToRemove.GetId() {
			customers[customersLength-1], customers[i] = customers[i], customers[customersLength-1]
			return customers[:customersLength-1]
		}
	}
	return customers
}
