package compiler

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"txtracker/pkg/logger"
)

// Compiler is an interface that defines the methods that a compiler must implement.
type Compiler interface {
	SolidityToAST_JSON(SolidityPath string) error
}

type SolidityCompiler struct {
}

// NewSolidityCompiler creates a new SolidityCompiler.

func NewSolidityCompiler() Compiler {
	return &SolidityCompiler{}
}

func (s *SolidityCompiler) SolidityToAST_JSON(SolidityPath string) error {
	fmt.Println("SolidityToAST_JSON called with path:", SolidityPath)

	// source virtual environment at ../../venv/bin/activate
	var cmdString string
	//fmt.Println(runtime.GOOS)
	if runtime.GOOS == "darwin" {
		fmt.Println("macOS detected")
		cmdString = "solc --ast-compact-json " + SolidityPath
	} else if runtime.GOOS == "linux" {
		fmt.Println("Linux detected")
		cmdString = "source ../../venv/bin/activate && solc --ast-compact-json " + SolidityPath
	} else {
		fmt.Println("OS not detected")
		logger.Fatal.Println("OS not detected. ")
		panic("OS not detected.")
	}

	cmd := exec.Command("zsh", "-c", cmdString)

	// remove .sol and add .ast.json
	astFileName := SolidityPath + ".ast.json"
	astFile, err := os.Create(astFileName)
	if err != nil {
		return fmt.Errorf("error creating ast file: %v", err)
	}
	defer astFile.Close()

	cmd.Stdout = astFile

	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running solc: %v", err)
	}

	return nil
}
