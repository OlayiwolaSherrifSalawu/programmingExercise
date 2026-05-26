package main

func removeZeros(si []int) []int {
	// two pointer approach
	indexZ := 0

	for i, num := range si {
		if num == 0 {
			indexZ = i
			break
		}
	}
	for i := indexZ + 1; i < len(si); i++ {
		num := si[indexZ]
		if si[i] > 0 {
			si[indexZ] = si[i]
			si[i] = num
			indexZ += 1
		}
	}
	return si
}
