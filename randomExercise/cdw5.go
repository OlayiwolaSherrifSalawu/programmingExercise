package main

import (
	"strconv"
	"strings"
)

func ranges(s []int) string {
	var result strings.Builder
	start := 0
	end := 0
	count := 0
	hasChanged := false
	toChangeSt := false
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
			continue
		}
		if thENd < sLen && s[thENd] == checkIf {
			if !toChangeSt {
				
			}
		}
	}
	return result.String()
}
