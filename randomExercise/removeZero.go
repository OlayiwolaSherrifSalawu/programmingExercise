package main

import "slices"

func removeZeros(si []int) []int {
	count := 0
	i := 0
	for {
		theO := si[i]
		if si[i] == 0 {
			si = slices.Delete(si, i, i+1)
			si = append(si, theO)
			count++
			i--
		}
		i++
		if count+i == len(si)-1 {
			break
		}
	}
	return si
}
