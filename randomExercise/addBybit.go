package main

func howShiftworks(a, b int) int {
	for b != 0 {
		a = a | b
		b = b >> 1
	}
	return a
}
