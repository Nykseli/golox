// +build gloxcompiler,!gloxvm,!gloxrun

// main file for compliler
// compailer turns source code to byte code that can be fed to
// virtual machine later

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// DefaultFileMod defines in what mode the compiled file will be by default
var DefaultFileMod os.FileMode = 0644

var vm = VM{}

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

func writeBytes(bytes []byte) {
	err := ioutil.WriteFile("out.glb", bytes, DefaultFileMod)

	if err != nil {
		log.Fatal("Creating glb error: ", err)
	}

}

func runFile(path string) {
	chunk := Chunk{}
	chunk.InitChunk()
	source := readFile(path)

	if !Compile(source, &chunk) {
		os.Exit(65)
	}

	var chunkBytes bytes.Buffer
	enc := gob.NewEncoder(&chunkBytes)

	err := enc.Encode(chunk)

	if err != nil {
		log.Fatal("encode error:", err)
	}

	writeBytes(chunkBytes.Bytes())

}

func mainTarget() {
	// Initialize vm
	vm.InitVM()

	if len(os.Args) == 1 {
		//repl()
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		fmt.Fprintf(os.Stderr, "Usage: gloxc [path]\n")
		os.Exit(64)
	}
}
