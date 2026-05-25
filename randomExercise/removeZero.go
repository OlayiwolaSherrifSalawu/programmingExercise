package main

import "slices"

func removeZeros(si []int) []int {
	for i := 0; i < len(si); i++ {
		theO := si[i]
		if si[i] == 0 {
			si = slices.Delete(si, i, i+1)
			si = append(si, theO)
		}
	}
	return si
}
