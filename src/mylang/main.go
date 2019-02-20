package main

// main.go purpose is to call the main function of the program its compiled to
// This file should only contain variables and functions that are used by
// every part

// This program is divided in three parts: runner, compiler and virtual machine
// Runner (main_runner.go) will compiles the source code and feeds it to virtual machine
//
// Compiler (main_compiler.go) will compile the source code and outputs the bytecode and data to out.glb
// glb stands for GLox Binary
//
// Virual machine ( main_vm.go) will read the out.glb, decode it and feeds it to virtual machine

// DebugPrintCode if true, prints dissasembled chunk from compiler
var DebugPrintCode = true

func main() {
	// Target files:
	// main_compiler.go (tag: gloxcompiler)
	// main_runner.go (tag: gloxrun)
	// main_vm.go (tag: gloxvm)
	mainTarget()
}
