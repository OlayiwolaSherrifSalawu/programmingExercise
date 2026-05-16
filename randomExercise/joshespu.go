package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(JospehPermutation([]int{1, 2, 3, 4, 5, 6, 7}, 3))
}
func JospehPermutation(s []int, kill int) []int {

	res := []int{}
	killCount := 1
	for len(s) > 0 {
		if kill > len(s) {
			if killCount == kill {
				res = append(res, (len(s)%killCount)-1)
				s = slices.Delete(s, (len(s)%killCount)-1, (len(s) % killCount))
				killCount = 1
			}
		}
		if killCount == kill {
			res = append(res, s[killCount-1])
			s = slices.Delete(s, killCount-1, killCount)
			killCount = 1
		}
		theVal := s[killCount-1]
		s = slices.Delete(s, killCount-1, killCount)
		s = append(s, theVal)
		killCount++
	}
	return res
}
