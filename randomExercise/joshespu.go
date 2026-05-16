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
		targetedPerson := s[0]
		if kill == killCount {
			res = append(res, s[0])
			s = slices.Delete(s, 0, 1)

			killCount = 1
		} else {
			s = slices.Delete(s, 0, 1)
			s = append(s, targetedPerson)
		}
		killCount++
	}
	return res
}
