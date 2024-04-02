package filehandler

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"txtracker/pkg/filehandler"
)

// Setup and Teardown functionality can be used to create a temporary testing environment.
func setupTestEnvironment() (cleanupFunc func(), DataPath string) {
	// Create temp directory
	tmpDir, err := os.MkdirTemp("", "contracts")
	if err != nil {
		panic(err)
	}
	// Create mock contract files
	os.WriteFile(filepath.Join(tmpDir, "TestContract1.sol"), []byte("contract TestContract1 {}"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "TestContract2.evm"), []byte("0x6001600081"), 0644)

	return func() { os.RemoveAll(tmpDir) }, tmpDir
}

func TestSolidityFileHandler_GetContractList(t *testing.T) {
	// Setup a test environment
	cleanup, DataPath := setupTestEnvironment()
	defer cleanup()

	handler := filehandler.NewSolidityFileHandler(DataPath)
	contractList, err := handler.GetContractList()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := []string{"TestContract1", "TestContract2"}
	if !reflect.DeepEqual(contractList, expected) {
		t.Errorf("Expected contract list to be %v, got %v", expected, contractList)
	}
}

func TestSolidityFileHandler_GetContractData(t *testing.T) {
	// Assume setupTestEnvironment has been called in a test suite setup if needed
	cleanup, DataPath := setupTestEnvironment()
	defer cleanup()

	handler := filehandler.NewSolidityFileHandler(DataPath)
	contractName := []string{"TestContract1", "TestContract2"}

	var solResults []string
	var evmResults []string
	for _, name := range contractName {
		sol, evm, err := handler.GetContractData(name)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		solResults = append(solResults, sol)
		evmResults = append(evmResults, evm)
	}

	sol_expected := []string{"contract TestContract1 {}", ""}
	evm_expected := []string{"", "0x6001600081"}
	if !reflect.DeepEqual(solResults, sol_expected) {
		t.Errorf("Expected contract data for %s, got %s", sol_expected, solResults)
	}
	if !reflect.DeepEqual(evmResults, evm_expected) {
		t.Errorf("Expected contract data for %s, got %s", evm_expected, evmResults)
	}
}
