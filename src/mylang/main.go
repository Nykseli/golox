package main

func main() {
	// Initialize vm
	vm := VM{}
	vm.InitVM()

	chunk := Chunk{}
	chunk.InitChunk()

	constant := chunk.AddConstant(1.2)
	chunk.WriteChunk(OpConstant, 123)
	chunk.WriteChunk(uint8(constant), 123)
	chunk.WriteChunk(OpNegate, 123)

	constant = chunk.AddConstant(3.4)
	chunk.WriteChunk(OpConstant, 123)
	chunk.WriteChunk(uint8(constant), 123)

	chunk.WriteChunk(OpAdd, 123)

	constant = chunk.AddConstant(5.6)
	chunk.WriteChunk(OpConstant, 123)
	chunk.WriteChunk(uint8(constant), 123)

	chunk.WriteChunk(OpDivide, 123)

	chunk.WriteChunk(OpReturn, 123)
	chunk.PrintHexes("Test Chunk")
	chunk.DisassembleChunk("Test Chunk")
	vm.Interpret(chunk)

}
