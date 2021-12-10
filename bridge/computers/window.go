package computers

import (
	"fmt"
	"structural-pattern-go/bridge/printers"
)

type windowComputer struct {
	printer printers.Printer
}

func NewWindowComputer() Computer {
	return &windowComputer{}
}

func (computer *windowComputer) Print() {
	fmt.Println("Print request for window")
	computer.printer.PrintFile()
}

func (computer *windowComputer) SetPrinter(printer printers.Printer) {
	computer.printer = printer
}
