package main

import "fmt"

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

func (vm *VM) readByte() uint8 {
	val := vm.IPArr[vm.IP]
	vm.IP++
	return val
}

func (vm *VM) binaryOp(op uint8) {
	b := vm.Pop()
	a := vm.Pop()

	switch op {
	case '+':
		vm.Push(a + b)
		break
	case '-':
		vm.Push(a - b)
		break
	case '*':
		vm.Push(a * b)
		break
	case '/':
		vm.Push(a / b)
		break
	}
}

func (vm *VM) readConstant() Value {
	return vm.Chunk.Constants.Values[vm.readByte()]
}

// run is where the actual bytecode is executed
func (vm *VM) run() int {
	for {
		instruction := vm.readByte()
		switch instruction {
		case OpConstant:
			{
				constant := vm.readConstant()
				vm.Push(constant)
				PrintValue(constant)
				fmt.Printf("\n")
				break
			}
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
		case OpNegate:
			vm.Push(-vm.Pop())
			break
		case OpReturn:
			PrintValue(vm.Pop())
			fmt.Printf("\n")
			return InterpretOk
		default:
			break

		}
	}
}

// Interpret chunk
func (vm *VM) Interpret(chunk Chunk) int {
	vm.Chunk = chunk
	vm.IP = 0
	vm.IPArr = vm.Chunk.Code
	return vm.run()
}
