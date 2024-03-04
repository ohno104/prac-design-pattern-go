package main

import "fmt"

type Monitor struct{}

func (m *Monitor) InputUsbcSource(conn Conn) {
	fmt.Println("input usbc source")
	conn.ConnUsbc()
}

type Conn interface {
	ConnUsbc()
}

type Usbc struct{}

func (u *Usbc) ConnUsbc() {
	fmt.Println("usbc port is connected")
}

type Hdmi struct{}

func (h *Hdmi) ConnHdmi() {
	fmt.Println("hdmi port is connected")
}

type HdmiAdapter struct {
	hdmi *Hdmi
}

func (h *HdmiAdapter) ConnUsbc() {
	fmt.Println("adapter modifies hdmi port to allowed connection")
	h.hdmi.ConnHdmi()
}

func main() {
	var _ Conn = (*Usbc)(nil) //!!!
	monitor := &Monitor{}
	usbc := &Usbc{}
	monitor.InputUsbcSource(usbc)

	hdmiAdapter := &HdmiAdapter{
		&Hdmi{},
	}

	monitor.InputUsbcSource(hdmiAdapter)
}
