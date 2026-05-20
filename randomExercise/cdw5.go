package main

import (
	"strconv"
	"strings"
)

func ranges(s []int) string {
	var result strings.Builder
	start := 0
	// end := 0
	count := 0
	hasChanged := false
	toChangeSt := true
	sLen := len(s)
	for i, val := range s {
		if hasChanged {
			hasChanged = false
			continue
		}
		checkIf := val + 1
		thENd := i + 1
		if count < 2 && thENd < sLen && s[thENd] != checkIf {
			result.WriteString(strconv.Itoa(s[thENd]))
			result.WriteString(",")
			count = 0
			continue
		}
		if thENd < sLen && s[thENd] == checkIf {
			if toChangeSt {
				start = i
				toChangeSt = false
			}
			count++
		}
		if thENd < sLen && s[thENd] != checkIf && count >= 2 {
			toChangeSt = true
			result.WriteString(strconv.Itoa(s[start]))
			result.WriteString("-")
			result.WriteString(strconv.Itoa(s[thENd-1]))
			result.WriteString(",")
			count = 0
		}
	}
	return result.String()
}
