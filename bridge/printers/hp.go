package printers

import "fmt"

type hpPrinter struct {
}

func NewHpPrinter() Printer {
	return &hpPrinter{}
}

func (printer *hpPrinter) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
