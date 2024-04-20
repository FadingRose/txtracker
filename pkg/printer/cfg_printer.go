package printer

import (
	"fmt"
	CFG "txtracker/pkg/cfg"
)

type CFGPrinter struct {
	CFG *CFG.CFG
}

func NewCFGPrinter(cfg *CFG.CFG) *CFGPrinter {
	return &CFGPrinter{
		CFG: cfg,
	}
}

func (p *CFGPrinter) Print() {
	for _, entry := range p.CFG.EntryPoints {
		fmt.Println("Entry Point -->", entry.Name)
		p.printFunction(entry)
		fmt.Println()
	}
}

func (p *CFGPrinter) printFunction(f *CFG.Function) {

	p.printBlock(f.Block)
}

func (p *CFGPrinter) printBlock(b *CFG.Block) {
	for _, s := range b.Statements {
		p.printStatement(s)
	}
}

func (p *CFGPrinter) printStatement(s *CFG.Statement) {
	tp := s.Type
	fmt.Println(tp.String())
}
