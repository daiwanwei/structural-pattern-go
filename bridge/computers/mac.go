package computers

import (
	"fmt"
	"structural-pattern-go/bridge/printers"
)

type macComputer struct {
	printer printers.Printer
}

func NewMacComputer() Computer {
	return &macComputer{}
}

func (computer *macComputer) Print() {
	fmt.Println("Print request for mac")
	computer.printer.PrintFile()
}

func (computer *macComputer) SetPrinter(printer printers.Printer) {
	computer.printer = printer
}
