package main

import (
	"fmt"
)

func main() {
	myMap := countAppearance([]string{"apple", "apple", "ola", "jide", "ola", "jimoh", "john", "done", "guy", "boy", "buy", "bone"})
	cleanedMap := myMap
	groups := groupLength(cleanedMap)
	fmt.Println(groups)
	finalSlice := checkAndPush(groups)

	fmt.Println(sortSlice(finalSlice))

}

//

func countAppearance(input []string) map[string]int {
	if input == nil {
		return nil
	}
	words := make(map[string]int)
	for _, val := range input {
		words[val] = words[val] + 1
	}
	return words
}

func groupLength(args map[string]int) map[int][]string {
	groups := make(map[int][]string)
	for key := range args {
		groups[len(key)] = append(groups[len(key)], key)
	}
	return groups
}
func checkAndPush(args map[int][]string) []string {
	nonDulicate := []string{}
	var count int
	for _, vals := range args {
		for _, word := range vals {
			isSafe := true
			for j := 0; j < len(vals); j++ {
				if word == vals[j] {
					continue
				}
				count = 0
				for k := 0; k < len(word); k++ {
					if word[k] != vals[j][k] {
						count++
					}
				}
				if count == 1 {
					isSafe = false
					break
				}
			}
			if isSafe {
				nonDulicate = append(nonDulicate, word)
			}
		}
	}
	return nonDulicate
}
func sortSlice(s []string) []string {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] >= s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}

		}
	}
	return s
}
