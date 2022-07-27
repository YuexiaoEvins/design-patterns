package Adapter

import "fmt"

type Computer interface {
	InsertIntoLightningPort()
}

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type WindowsAdapter struct {
	windows *Windows
}

func (wa *WindowsAdapter) InsertIntoLightningPort() {
	// convert something
	fmt.Println("Adapter converts Lightning signal to USB.")
	wa.windows.InsertIntoUSBPort()
}

type Windows struct{}

func (w *Windows) InsertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}
