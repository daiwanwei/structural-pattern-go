package main

import "fmt"

type MacAdapter struct {
	Mac Mac
}

func (adapter *MacAdapter) ConnectWithHdmi() {
	fmt.Println("adapter of Hdmi and typeC")
	adapter.Mac.ConnectWithTypeC()
}

func (adapter *MacAdapter) ConnectWithDvi() {
	fmt.Println("adapter of DVI and typeC")
	adapter.Mac.ConnectWithTypeC()
}
