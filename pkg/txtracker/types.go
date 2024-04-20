package txtracker

import CFG "txtracker/pkg/cfg"

type TxSeQuence struct {
	Name string
	Tx   []Tx
}

type Tx struct {
	Name       string
	Statements []CFG.Statement
}
