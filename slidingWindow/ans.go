package main

import (
	"fmt"
)

func main() {
	// fmt.Println(maxTemp([]int{70, 75, 80, 72, 74, 90, 85, 88, 86, 80, 60}, 3))
	// fmt.Println(correctOrder("act", "cat"))
	fmt.Println(stripHtml("<b>bold</b> <html>ola</html>"))
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
		"<":  ">",
		"</": ">",
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
