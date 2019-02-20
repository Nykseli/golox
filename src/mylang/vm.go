package main

import (
	"fmt"
	"os"
)

// RunTimeError tells if vm has encountered an error
var RunTimeError = false

// StackMax defines the maximum size of the VM stack
const StackMax = 256

const (
	// InterpretOk is returned when program ran succesfully
	InterpretOk = iota
	// InterpretCompileError is returned when compiling failed
	InterpretCompileError = iota
	// InterpretRuntimeError is returned when running the bytecode failed
	InterpretRuntimeError = iota
)

// VM is virtual mashine that runs the bytecode
type VM struct {
	Chunk    Chunk
	IP       uint8
	IPArr    []uint8
	Stack    [StackMax]Value
	StackTop Value
	// StackPos keeps track of the stack position
	StackPos int
}

func (vm *VM) resetStack() {
	vm.StackTop = vm.Stack[vm.StackPos]
}

func runTimeError(format string, args ...string) {
	fmt.Fprintf(os.Stderr, format)
	fmt.Fprintf(os.Stderr, "\n")

	fmt.Fprintf(os.Stderr, "[line %d] in script\n",
		vm.Chunk.Lines[vm.IP])

}

// InitVM initializes the virtual mashine
func (vm *VM) InitVM() {
	vm.resetStack()
}

// FreeVM frees the VM state
func (vm *VM) FreeVM() {
	//TODO:
}

// Push Value to stack
func (vm *VM) Push(value Value) {
	vm.Stack[vm.StackPos] = value
	vm.StackPos++
}

// Pop Value from stack
func (vm *VM) Pop() Value {
	vm.StackPos--
	return vm.Stack[vm.StackPos]
}

func (vm *VM) peekStack(distance int) Value {
	return vm.Stack[vm.StackPos-1-distance]
}

func isFalsey(value Value) bool {
	return IsNil(value) || (IsBool(value) && !AsBool(value))
}

func (vm *VM) readByte() uint8 {
	val := vm.IPArr[vm.IP]
	vm.IP++
	return val
}

func (vm *VM) binaryOp(op uint8) {

	if !IsNumber(vm.peekStack(0)) || !IsNumber(vm.peekStack(1)) {
		runTimeError("Operands must be numbers.")
		RunTimeError = true
		return
	}

	b := AsNumber(vm.Pop())
	a := AsNumber(vm.Pop())

	switch op {
	case '<':
		vm.Push(BoolVal(a < b))
		break
	case '>':
		vm.Push(BoolVal(a > b))
		break
	case '+':
		vm.Push(NumberVal(a + b))
		break
	case '-':
		vm.Push(NumberVal(a - b))
		break
	case '*':
		vm.Push(NumberVal(a * b))
		break
	case '/':
		vm.Push(NumberVal(a / b))
		break
	}
}

func (vm *VM) readConstant() Value {
	return vm.Chunk.Constants.Values[vm.readByte()]
}

// run is where the actual bytecode is executed
func (vm *VM) run() int {
	for {
		fmt.Printf("          ")
		for i := 0; i < int(vm.StackPos); i++ {
			if vm.Stack[i].As != nil {
				fmt.Printf("[ ")
				PrintValue(vm.Stack[i])
				fmt.Printf(" ]")
			}
		}
		fmt.Printf("\n")
		vm.Chunk.DisassembleInstruction(int(vm.IP))

		instruction := vm.readByte()
		switch instruction {
		case OpConstant:
			{
				constant := vm.readConstant()
				vm.Push(constant)
				break
			}
		case OpNil:
			vm.Push(NilVal())
			break
		case OpTrue:
			vm.Push(BoolVal(true))
			break
		case OpFalse:
			vm.Push(BoolVal(false))
			break
		case OpEqual:
			{
				b := vm.Pop()
				a := vm.Pop()
				vm.Push(BoolVal(ValuesEqual(a, b)))
				break
			}
		case OpGreater:
			vm.binaryOp('>')
			break
		case OpLess:
			vm.binaryOp('<')
			break
		case OpAdd:
			vm.binaryOp('+')
			break
		case OpSubtract:
			vm.binaryOp('-')
			break
		case OpMultiply:
			vm.binaryOp('*')
			break
		case OpDivide:
			vm.binaryOp('/')
			break
		case OpNot:
			vm.Push(BoolVal(isFalsey(vm.Pop())))
		case OpNegate:
			if !IsNumber(vm.peekStack(0)) {
				runTimeError("Operand must be a number.")
				RunTimeError = true
			}
			vm.Push(NumberVal(-AsNumber(vm.Pop())))
			break
		case OpReturn:
			PrintValue(vm.Pop())
			fmt.Printf("\n")
			return InterpretOk
		default:
			break

		}

		if RunTimeError {
			return InterpretRuntimeError
		}
	}
}

// InterpretBytes feeds the chunk that we get from glb file
func (vm *VM) InterpretBytes(chunk Chunk) int {
	vm.Chunk = chunk
	vm.IP = 0
	vm.IPArr = vm.Chunk.Code
	return vm.run()
}

// Interpret from source string
func (vm *VM) Interpret(source string) int {
	var chunk = Chunk{}
	chunk.InitChunk()
	if !Compile(source, &chunk) {
		return InterpretCompileError
	}

	vm.Chunk = chunk
	vm.IP = 0
	vm.IPArr = vm.Chunk.Code
	return vm.run()
}
