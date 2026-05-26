package main

import (
	"slices"
)

func removeZeros(si []int) []int {
	// two pointer approach
	i := slices.Index(si, 0)
	if i < 0 {
		return si
	}
	for j := i + 1; j < len(si); j++ {
		if si[j] > 0 {
			temp := si[i]
			si[i] = si[j]
			si[j] = temp
			i++
		}
	}
	return si

}
