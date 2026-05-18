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
	for i, val := range s {
		checkIf := val + 1
		if i+1 < len(s) && s[i+1] == checkIf {
			if !hasChanged {
				start = i
				hasChanged = true
			}
			count++

		}
		if count == 2 && hasChanged {
			end = i + 1
			result.WriteString(strconv.Itoa(s[start]))
			result.WriteString("-")
			result.WriteString(strconv.Itoa(s[end]))
			result.WriteString(",")
			hasChanged = false
			continue
		}
		if i+1 < len(s) && s[i+1] != checkIf && count > 0 {
			end = i + 1
			result.WriteString(strconv.Itoa(s[end]))
			result.WriteString(",")
			count = 0

			continue
		}

	}
	return result.String()
}
