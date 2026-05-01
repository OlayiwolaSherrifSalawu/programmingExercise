package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println(maxTemp([]int{70, 75, 80, 72, 74, 90, 85, 88, 86, 80, 60}, 3))
	// fmt.Println(correctOrder("act", "cat"))
	fmt.Printf("%q\n", AddSpaceAndSplit("Hello,"))
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

func stripHtml(s string) string {
	tag := map[string]string{
		"<": ">",
	}
	store := ""
	tags := ""
	foundTag := false
	for i := 0; i < len(s); i++ {
		val, ok := tag[string(s[i])]
		if ok {
			tags = val

		}
		if ok {
			foundTag = true
			continue
		}
		if string(s[i]) == tags {
			foundTag = false
			continue
		}
		if !ok && !foundTag {
			store += string(s[i])
		}
		if !ok && foundTag {
			continue
		}
	}
	return store
}

func dnaAnalyser(s string, n int) string {
	if len(s) < n {
		return ""
	}
	maxCount := 0
	sequce := ""
	maxSequence := ""
	vowels := "aeiouAEIOU"
	count := 0
	for i := 0; i < n; i++ {
		if strings.Contains(vowels, string(s[i])) {
			maxCount++
		}
		sequce += string(s[i])
	}
	maxSequence = sequce
	count = maxCount

	for j := n; j < len(s); j++ {
		sequce = sequce[1:] + string(s[j])
		if strings.Contains(vowels, string(s[j])) {
			count++
		}
		if strings.Contains(vowels, string(s[j-n])) {
			count--
		}
		if count > maxCount {
			maxCount = count
			maxSequence = sequce
		}
	}
	return maxSequence
}

func AddSpaceAndSplit(s string) []string {
	res := ""
	for _, val := range s {
		if strings.Contains(",.:;!?", string(val)) {
			res += " " + string(val)
			continue
		}
		res += string(val)
	}
	resul := strings.Fields(res)

	return resul
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	revStore := "" //because we wanted to reverse string
	store := ""
	if n < 0 {
		n = -n
		revStore = "-"
	}
	// itoa core functions (loop)
	for n != 0 {
		val := n % 10
		store += string('0' + val)
		n = n / 10
	}
	for i := len(store) - 1; i >= 0; i-- {
		revStore += string(store[i])
	}

	return store
}
func Union(s1, s2 string) string {
	newString := s1 + s2
	union := ""
	for i := 0; i < len(newString); i++ {
		if !strings.ContainsAny(union, string(newString[i])) {
			union += string(newString[i])
		}
	}
	return union
}

func FithAndSkip(s string) string {
	if len(s) < 5 {
		return "Invalid string"
	}
	var sb strings.Builder
	count := 0
	for _, val := range s {
		if val == ' ' {
			continue
		}
		if count == 5 {
			sb.WriteRune(' ')
			count = 0
		}
		sb.WriteRune(val)
		count++
	}
	return sb.String()
}

func NotDecimal(s string) string {
	if s == "" {
		return s
	}
	oldVal := 0
	minus := ""
	if s[0] == '-' {
		minus = "-"
		s = s[1:]
	}
	for _, val := range s {
		if val == '.' {
			continue
		}
		if !(val >= '0' && val <= '9') {
			return minus + s
		}
		cal := 10*oldVal + int(val-'0')
		oldVal = cal
	}
	result := strconv.Itoa(oldVal)
	return minus + result
}
