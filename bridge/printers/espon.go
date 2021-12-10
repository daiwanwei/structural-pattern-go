package printers

import "fmt"

type epsonPrinter struct {
}

func NewEpsonPrinter() Printer {
	return &epsonPrinter{}
}

func (printer *epsonPrinter) PrintFile() {
	fmt.Println("Printing by a epson Printer")
}
