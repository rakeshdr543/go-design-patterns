package adapter

import "fmt"

type Mac struct {
}

func (m *Mac) ConnectToTypeC() {
	fmt.Println("Mac connected to type C")
}
