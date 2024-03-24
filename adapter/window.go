package adapter

type Window struct {
}

func (w *Window) ConnectToUSB() {
	println("Window connected to USB")
}
