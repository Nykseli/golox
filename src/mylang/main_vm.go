// +build gloxvm,!gloxrun,!gloxcompiler

// Main file for glox virual machine
// virual machine loads chucks from binary file and feeds runs it with vm

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var vm = VM{}

// readFile into []byte
func readFile(path string) []byte {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(74)
	}

	return fileBytes
}

// run file get bytes from path with readFileFunction
// decodes it to struct and feeds it to vm
func runFile(path string) {
	chunk := Chunk{}
	chunk.InitChunk()
	source := readFile(path)

	chunkBytes := bytes.NewReader(source)

	dec := gob.NewDecoder(chunkBytes)

	var chunkStruct Chunk
	err := dec.Decode(&chunkStruct)

	if err != nil {
		log.Fatal("decode error:", err)
	}

	vm.InterpretBytes(chunkStruct)

}

func mainTarget() {
	// Initialize vm
	vm.InitVM()

	if len(os.Args) == 1 {
		//repl()
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		fmt.Fprintf(os.Stderr, "Usage: gloxvm [path]\n")
		os.Exit(64)
	}
}
