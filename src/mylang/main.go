// +build gloxrun,!gloxvm,!gloxcompiler

// main file for building runner
// runner parses the source code to bytecode and feeds

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// DebugPrintCode if true, prints dissasembled chunk from compiler
var DebugPrintCode = true

var vm = VM{}

func repl() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")

		line, err := reader.ReadString('\n')

		if err != nil || len(line) == 0 {
			fmt.Printf("\n")
			break
		}

		// reader.Readstring doesn't include the the EOF byte so we need to add it
		line += string(FileEOF)

		vm.Interpret(line)
	}
}

func readFile(path string) string {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(74)
	}

	// ioutil.ReadFile doesn't include the the EOF byte so we need to add it
	sourceBytes := append(fileBytes, FileEOF)

	return string(sourceBytes)
}

func runFile(path string) {
	source := readFile(path)
	result := vm.Interpret(source)

	if result == InterpretCompileError {
		os.Exit(65)
	}
	if result == InterpretRuntimeError {
		os.Exit(70)
	}
}

func main() {
	// Initialize vm
	vm.InitVM()

	if len(os.Args) == 1 {
		repl()
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		fmt.Fprintf(os.Stderr, "Usage: gloxrun [path]\n")
		os.Exit(64)
	}
}
