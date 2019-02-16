package main

import (
	"fmt"
)

// DisassembleChunk prints the chunk in human readable form
func (chunk *Chunk) DisassembleChunk(name string) {
	fmt.Printf("== %s == \n", name)

	for offset := 0; offset < chunk.Count; {
		offset = chunk.DisassembleInstruction(offset)
	}
}

// DisassembleInstruction prints the instruction in chunk in human readable form
func (chunk *Chunk) DisassembleInstruction(offset int) int {
	fmt.Printf("%04d ", offset)

	if offset > 0 && chunk.Lines[offset] == chunk.Lines[offset-1] {
		fmt.Printf("   | ")
	} else {
		fmt.Printf("%4d ", chunk.Lines[offset])
	}

	instruction := chunk.Code[offset]
	switch instruction {
	case OpConstant:
		return chunk.constantInstruction("OP_CONSTANT", offset)
	case OpReturn:
		return chunk.simpleInstruction("OP_RETURN", offset)
	default:
		fmt.Printf("Unknown upcode %d\n", instruction)
		return offset + 1

	}
}

func (chunk *Chunk) constantInstruction(name string, offset int) int {
	constant := chunk.Code[offset+1]
	fmt.Printf("%-16s %4d '", name, constant)
	PrintValue(chunk.Constants.Values[constant])
	fmt.Printf("'\n")
	return offset + 2
}

func (chunk *Chunk) simpleInstruction(name string, offset int) int {
	fmt.Printf("%s\n", name)
	return offset + 1
}
