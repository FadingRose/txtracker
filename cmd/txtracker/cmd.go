package main

import "os"

type PrinterType string

const (
	CFG_PRINTER       PrinterType = "cfg"
	CALLGRAPH_PRINTER PrinterType = "callgraph"
)

type SPECIFIC_CONTRACT = string

func cmd(args []string) (SPECIFIC_CONTRACT, PrinterType) {
	var SPECIFIC_CONTRACT string
	var PRINTER PrinterType
	SPECIFIC_CONTRACT, PRINTER = "", CFG_PRINTER
	if len(os.Args) >= 2 {
		SPECIFIC_CONTRACT = os.Args[1]
	} else {
		SPECIFIC_CONTRACT = ""
	}
	return PRINTER, SPECIFIC_CONTRACT

}
