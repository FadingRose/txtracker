package ast

import (
	"reflect"
	"testing"
	"txtracker/pkg/ast"
)

func TestSourceUnitConstructor(t *testing.T) {
	data := map[string]interface{}{
		"absolutePath":         "../../dataset/contracts/0x0a3f9678d6b631386c2dd3de8809b48b0d1bbd56.sol",
		"experimentalSolidity": true,
		"exportedSymbols": map[string][]float64{
			"ERC20":       {275},
			"LikerCoin":   {1738},
			"LockBalance": {1554},
			"Ownable":     {151},
			"Pausable":    {208},
			"SafeMath":    {97},
			"Token":       {847},
		},
		"license": "MIT",
	}

	expected := ast.SourceUnit{
		AbsolutePath:         "../../dataset/contracts/0x0a3f9678d6b631386c2dd3de8809b48b0d1bbd56.sol",
		ExperimentalSolidity: true,
		ExportedSymbols: map[string][]int{
			"ERC20":       {275},
			"LikerCoin":   {1738},
			"LockBalance": {1554},
			"Ownable":     {151},
			"Pausable":    {208},
			"SafeMath":    {97},
			"Token":       {847},
		},
		License: "MIT",
	}

	var su ast.SourceUnit
	su.Constructor(&data)

	if !reflect.DeepEqual(su, expected) {
		t.Errorf("Constructor() = %v, want %v", su, expected)
	}
}

func TestSourceUnitAttributes(t *testing.T) {
	su := ast.SourceUnit{
		AbsolutePath:         "../../dataset/contracts/0x0a3f9678d6b631386c2dd3de8809b48b0d1bbd56.sol",
		ExperimentalSolidity: true,
		ExportedSymbols: map[string][]int{
			"ERC20":       []int{275},
			"LikerCoin":   []int{1738},
			"LockBalance": []int{1554},
			"Ownable":     []int{151},
			"Pausable":    []int{208},
			"SafeMath":    []int{97},
			"Token":       []int{847},
		},
		License: "MIT",
	}

	attrs := su.Attributes()

	expectedAttrs := map[string]interface{}{
		"AbsolutePath":         "../../dataset/contracts/0x0a3f9678d6b631386c2dd3de8809b48b0d1bbd56.sol",
		"ExperimentalSolidity": true,
		"ExportedSymbols": map[string][]int{
			"ERC20":       []int{275},
			"LikerCoin":   []int{1738},
			"LockBalance": []int{1554},
			"Ownable":     []int{151},
			"Pausable":    []int{208},
			"SafeMath":    []int{97},
			"Token":       []int{847},
		},
		"License": "MIT",
	}

	if !reflect.DeepEqual(*attrs, expectedAttrs) {
		t.Errorf("Attributes() = %v, want %v", *attrs, expectedAttrs)
	}
}
