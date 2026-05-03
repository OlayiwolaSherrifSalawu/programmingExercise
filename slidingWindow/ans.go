package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(maxTemp([]int{70, 75, 80, 72, 74, 90, 85, 88, 86, 80, 60}, 3))
	// fmt.Println(correctOrder("act", "cat"))
	fmt.Println(dnaAnalyser("xyaweio4tuidgertfgieieieu", 6))
}

func maxTemp(slic []int, n int) ([]int, int) {
	var max int
	var index int
	var oldTemp int
	for i := 0; i < n; i++ {
		oldTemp += slic[i]
	}

	for i, j := n, 0; i < len(slic); i++ {
		newTemp := oldTemp + slic[i] - slic[j]
		if newTemp > oldTemp {
			max = newTemp
			index = i
		}
		oldTemp = newTemp
		j++

	}
	if max == oldTemp {
		return slic[:n+1], oldTemp
	}
	return slic[index-n+1 : index+1], max
}

func dnaAnalyser(s string, n int) string {
	if len(s) < n {
		return ""
	}
	maxCount := 0
	sequce := ""
	maxSequence := 0
	vowels := "aeiouAEIOU"
	count := 0
	for i := 0; i < n; i++ {
		if strings.Contains(vowels, string(s[i])) {
			maxCount++
		}
		sequce += string(s[i])
		maxSequence++
	}

	count = maxCount

	for j := n; j < len(s); j++ {

		if strings.Contains(vowels, string(s[j])) {
			count++
		}
		if strings.Contains(vowels, string(s[j-n])) {
			count--
		}
		if count > maxCount {
			maxCount = count
			maxSequence = j
		}
	}
	return s[maxSequence-n+1 : maxSequence+1]
}
func AnagramDective(s, target string) []int {
	isAnagram := true
	angramIndex := []int{}
	targetMap := make(map[string]int)
	wordMap := make(map[string]int)
	for _, val := range target {
		targetMap[string(val)] = targetMap[string(val)] + 1
	}
	for i := 0; i < len(target); i++ {
		wordMap[string(s[i])] = wordMap[string(s[i])] + 1
		_, targetVal := targetMap[string(s[i])]
		_, wordVal := wordMap[string(s[i])]
		if targetVal != wordVal {
			isAnagram = false
		}
	}
	if isAnagram {
		angramIndex = append(angramIndex, 0)
	}
}
