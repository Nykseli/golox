package main

// GrowCapacity doubles the oldCapacity
func GrowCapacity(oldCapacity int) int {
	if oldCapacity < 8 {
		oldCapacity = 8
	} else {
		oldCapacity = (oldCapacity * 2)
	}

	return oldCapacity
}
