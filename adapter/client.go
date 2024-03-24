package adapter

import "fmt"

type Client struct {
}

func (c *Client) ConnectToTypeC(comp IComputer) {
	fmt.Println("Client connected to type C")
	comp.ConnectToTypeC()
}
