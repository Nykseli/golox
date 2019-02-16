package main

// OpCode is for OpCode "enum"
type OpCode uint8

const (
	// OpConstant is code for constant value
	OpConstant uint8 = iota // 0
	// OpReturn is code for return
	OpReturn uint8 = iota // 1
)

// Chunk contains the program code in bytecodes
type Chunk struct {
	Count     int
	Capacity  int
	Code      []uint8
	Lines     []int
	Constants ValueArray
}

// InitChunk sets the initial values
func (chunk *Chunk) InitChunk() {
	chunk.Count = 0
	chunk.Capacity = 0
	chunk.Code = nil
	chunk.Lines = nil
	chunk.Constants = ValueArray{}
}

// FreeChunk just initializes the object and lets the golang deal with memory
func (chunk *Chunk) FreeChunk() {
	chunk.InitChunk()
}

// WriteChunk writes instruction to chunk struct
func (chunk *Chunk) WriteChunk(_byte uint8, line int) {
	if chunk.Capacity < chunk.Count+1 {
		oldCapacity := chunk.Capacity
		chunk.Capacity = GrowCapacity(oldCapacity)
		chunk.Reallocate(oldCapacity)
	}

	chunk.Code[chunk.Count] = _byte
	chunk.Lines[chunk.Count] = line
	chunk.Count++
}

// AddConstant adds constant to chunk ValueArray
func (chunk *Chunk) AddConstant(value Value) int {
	chunk.Constants.WriteValueArray(value)
	return chunk.Constants.Count - 1
}

// Reallocate the Chunk
func (chunk *Chunk) Reallocate(oldSize int) {
	// Resize the Code
	t := make([]uint8, chunk.Capacity)
	copy(t, chunk.Code)
	chunk.Code = t

	// Resize the lines
	t1 := make([]int, chunk.Capacity)
	copy(t1, chunk.Lines)
	chunk.Lines = t1
}