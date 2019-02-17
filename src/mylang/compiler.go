package main

import "fmt"

// Compile the source code
func Compile(source string) {
	InitScanner(source)
	line := -1
	var token Token
	for {
		token = ScanToken()
		if token.Line != line {
			fmt.Printf("%4d ", token.Line)
			line = token.Line
		} else {
			fmt.Printf("   | ")
		}
		fmt.Printf("%2d '%s'\n", token.Type, token.Value)

		if token.Type == TokenEOF {
			break
		}

	}

}
