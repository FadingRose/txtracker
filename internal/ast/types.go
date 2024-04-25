package ast

type ExportedSymbols map[string][]int

// Constructors
func (e *ExportedSymbols) Constructor(data *map[string][]float64) {
	for key, value := range *data {
		(*e)[key] = make([]int, len(value))
		for i, v := range value {
			(*e)[key][i] = int(v)
		}
	}
}

type InternalFunctionIDs map[string]int

// Constructors
func (i *InternalFunctionIDs) Constructor(data *map[string]float64) {
	for key, value := range *data {
		(*i)[key] = int(value)
	}
}

type LinearizedBaseContracts []int

// Constructors
func (l *LinearizedBaseContracts) Constructor(data *[]float64) {
	for _, value := range *data {
		*l = append(*l, int(value))
	}
}

type ContractDependencies []int

func (c *ContractDependencies) Constructor(data *[]float64) {
	for _, value := range *data {
		*c = append(*c, int(value))
	}
}

type UsedErrors []int

func (u *UsedErrors) Constructor(data *[]float64) {
	for _, value := range *data {
		*u = append(*u, int(value))
	}
}

type UsedEvents []int

func (u *UsedEvents) Constructor(data *[]float64) {
	for _, value := range *data {
		*u = append(*u, int(value))
	}
}

type NameLocations []string

func (n *NameLocations) Constructor(data *[]string) {
	for _, value := range *data {
		*n = append(*n, value)
	}
}

type BaseFunctions []int

func (b *BaseFunctions) Constructor(data *[]float64) {
	for _, value := range *data {
		*b = append(*b, int(value))
	}
}

type Literals []string

func (l *Literals) Constructor(data *[]string) {
	for _, value := range *data {
		*l = append(*l, value)
	}
}
