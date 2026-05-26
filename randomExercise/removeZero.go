package main

func removeZeros(si []int) []int {
	// two pointer approach
	seenZero := false
	i := 0
	for j := 0; j < len(si); j++ {
		if si[i] == 0 && si[j] > 0 && seenZero {
			temp := si[i]
			si[i] = si[j]
			si[j] = temp
			i++
		}

		if si[i] == 0 && seenZero == false {
			seenZero = true
		}
		if si[i] != 0 && seenZero == false {
			i++
		}
	}
	return si

}
