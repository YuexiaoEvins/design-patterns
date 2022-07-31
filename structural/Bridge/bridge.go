package Bridge

import "fmt"

type Printer interface {
	PrintFile()
}

type Computer interface {
	Print()
	SetPrinter(Printer)
}

type Mac struct {
	printer Printer
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

func (m *Mac) Print() {
	fmt.Println("mac printing!")
	m.printer.PrintFile()
}

type Window struct {
	printer Printer
}

func (m *Window) SetPrinter(p Printer) {
	m.printer = p
}

func (m *Window) Print() {
	fmt.Println("Window printing!")
	m.printer.PrintFile()
}

type Epson struct {
}

func (e *Epson) PrintFile() {
	fmt.Println("Epson printer print!")
}

type HP struct {
}

func (e *HP) PrintFile() {
	fmt.Println("HP printer print!")
}
