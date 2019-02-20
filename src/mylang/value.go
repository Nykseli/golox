package main

import "fmt"

// ValueType defines how the Value is handeled
type ValueType uint8

const (
	// ValBool is type for true and false
	ValBool ValueType = iota
	// ValNil is type for nil
	ValNil ValueType = iota
	// ValNumber is type for all number
	ValNumber ValueType = iota
)

// BoolValue is for true or false
type BoolValue struct {
	Boolean bool
}

// NilValue is represented as 0 float64
type NilValue struct {
	Number float64
}

// NumberValue is numbers in float64
type NumberValue struct {
	Number float64
}

// Value is value for constants
type Value struct {
	Type ValueType
	As   interface{}
}

// IsBool checks if the value type is ValBool
func IsBool(value Value) bool {
	return value.Type == ValBool
}

// IsNil checks if the value type is ValNil
func IsNil(value Value) bool {
	return value.Type == ValNil
}

// IsNumber checks if the value type is ValNumber
func IsNumber(value Value) bool {
	return value.Type == ValNumber
}

// AsBool gets the boolean from the value
func AsBool(value Value) bool {
	return value.As.(BoolValue).Boolean

}

// AsNumber gets the float from the value
func AsNumber(value Value) float64 {
	return value.As.(NumberValue).Number
}

// BoolVal creates Value struct with ValBool type based on the value parameter
func BoolVal(value bool) Value {
	val := Value{}
	val.Type = ValBool
	val.As = BoolValue{value}

	return val
}

// NilVal creates Value struct with ValNil type
func NilVal() Value {
	val := Value{}
	val.Type = ValNil
	val.As = NilValue{0}

	return val
}

// NumberVal creates Value struct with ValNumber type based on the value parameter
func NumberVal(value float64) Value {
	val := Value{}
	val.Type = ValNumber
	val.As = NumberValue{value}

	return val

}

// ValueArray holds values
type ValueArray struct {
	Capacity int
	Count    int
	Values   []Value
}

// ValuesEqual cheks if values equal
// Equality between different types is always false
func ValuesEqual(a Value, b Value) bool {
	if a.Type != b.Type {
		return false
	}

	switch a.Type {
	case ValBool:
		return AsBool(a) == AsBool(b)
	case ValNil:
		return true
	case ValNumber:
		return AsNumber(a) == AsNumber(b)

	default:
		return false
	}
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
	switch value.Type {
	case ValBool:
		if AsBool(value) {
			fmt.Printf("true")
		} else {
			fmt.Printf("false")

		}
	case ValNil:
		fmt.Printf("nil")
	case ValNumber:
		fmt.Printf("%g", AsNumber(value))
	}
}
