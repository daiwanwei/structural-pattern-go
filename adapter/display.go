package main

import "fmt"

type DellDisplay struct {
}

func (display *DellDisplay) Connect(port DellPort) {
	fmt.Println("connect to display with hdmi")
	port.ConnectWithHdmi()
}

type DellPort interface {
	ConnectWithHdmi()
}
