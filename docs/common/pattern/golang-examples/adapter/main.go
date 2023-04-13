package main

import "fmt"

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct{}

func (m *MyLegacyPrinter) Print(s string) string {
	return fmt.Sprintf("Legacy Printer: %s\n", s)
}

type NewPrinter interface {
	PrintStored() string
}

type MyNewPrinter struct {
	data string
}

func (m *MyNewPrinter) PrintStored() string {
	return fmt.Sprintf("New Printer: %s\n", m.data)
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	NewPrinter *MyNewPrinter
}

func (p *PrinterAdapter) PrintStored() string {
	if p.NewPrinter == nil {
		p.NewPrinter = &MyNewPrinter{data: p.OldPrinter.Print("Adapter")}
	}
	return p.NewPrinter.PrintStored()
}

func main() {
	var adapter NewPrinter = &PrinterAdapter{
		OldPrinter: &MyLegacyPrinter{},
	}

	fmt.Println(adapter.PrintStored())
}
