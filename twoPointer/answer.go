package twopointer

func correctOrder(word, target string) bool {
	i := 0
	for j := 0; j < len(word); j++ {
		if target[i] == word[j] {
			i++
		}
	}

	if i == len(target) {
		return true
	}
	return false
}
