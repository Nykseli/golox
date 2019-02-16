package main

import "fmt"

// Value is value for constants
type Value float64

// ValueArray holds values
type ValueArray struct {
	Capacity int
	Count    int
	Values   []Value
}

// InitValueArray initializes the Value array
func (array *ValueArray) InitValueArray() {
	array.Values = nil
	array.Capacity = 0
	array.Count = 0
}

// WriteValueArray adds new alue to ValueArray
func (array *ValueArray) WriteValueArray(value Value) {
	if array.Capacity < array.Count+1 {
		oldCapacity := array.Capacity
		array.Capacity = GrowCapacity(oldCapacity)
		array.Reallocate(oldCapacity)
	}

	array.Values[array.Count] = value
	array.Count++
}

// Reallocate the ValueArray
func (array *ValueArray) Reallocate(oldSize int) {
	// Resize the Values
	t := make([]Value, array.Capacity)
	copy(t, array.Values)
	array.Values = t
}

// PrintValue prints the value
func PrintValue(value Value) {
	fmt.Printf("%g", value)
}
