package txtracker

type TxSeQuence struct {
	Name string
	Tx   []Tx
}

type Tx struct {
	Name       string
	Statements []Statement
}
