package adapter

import "fmt"

type WindowAdapter struct {
	Window *Window
}

func (wa *WindowAdapter) ConnectToTypeC() {
	fmt.Println("WindowAdapter connected to type C")
	wa.Window.ConnectToUSB()
}
