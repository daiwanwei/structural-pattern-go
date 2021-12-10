package main

import (
	"structural-pattern-go/bridge/computers"
	"structural-pattern-go/bridge/printers"
)

func main() {
	mac := computers.NewMacComputer()
	window := computers.NewWindowComputer()
	hp := printers.NewHpPrinter()
	espon := printers.NewEpsonPrinter()
	mac.SetPrinter(hp)
	window.SetPrinter(espon)
	mac.Print()
	window.Print()
}
