package computers

import "structural-pattern-go/bridge/printers"

type Computer interface {
	Print()
	SetPrinter(printer printers.Printer)
}
