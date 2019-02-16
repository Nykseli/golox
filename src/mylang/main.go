package main

func main() {
	chunk := Chunk{}
	chunk.InitChunk()

	constant := chunk.AddConstant(1.2)
	chunk.WriteChunk(OpConstant, 123)
	chunk.WriteChunk(uint8(constant), 123)

	chunk.WriteChunk(OpReturn, 123)
	chunk.DisassembleChunk("Test Chunk")

}
