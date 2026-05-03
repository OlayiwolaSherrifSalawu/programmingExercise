package main

import (
	"fmt"
	"maps"
	"strings"
)

func main() {
	// fmt.Println(maxTemp([]int{70, 75, 80, 72, 74, 90, 85, 88, 86, 80, 60}, 3))
	// fmt.Println(correctOrder("act", "cat"))
	fmt.Println(AnagramDective("cbaebabacd", "abc"))
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
	// "cbaebabacd", "abc"
	for i := 0; i < len(target); i++ {
		wordMap[string(s[i])] = wordMap[string(s[i])] + 1
		if targetMap[string(s[i])] != wordMap[string(s[i])] {
			isAnagram = false
		}
	}
	if isAnagram {
		angramIndex = append(angramIndex, 0)
	}
	for i := len(target); i < len(s); i++ {
		wordMap[string(s[i])] = wordMap[string(s[i])] + 1
		wordMap[string(s[i-len(target)])] = wordMap[string(s[i-len(target)])] - 1
		if wordMap[string(s[i-len(target)])] == 0 {
			delete(wordMap, string(s[i-len(target)]))
		}
		if maps.Equal(wordMap, targetMap) {
			angramIndex = append(angramIndex, i-len(target)+1)
		}
	}
	return angramIndex
}
