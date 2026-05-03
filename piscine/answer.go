package main

import (
	"slices"
	"strconv"
	"strings"
)

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

func Brackets(s string) bool {
	brackets := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	stack := []string{}
	for _, val := range s {
		sVal := string(val)
		if vals, ok := brackets[sVal]; ok {
			stack = append(stack, vals)
		}
		if len(stack) != 0 {
			if vals, _ := brackets[stack[len(stack)-1]]; vals == sVal {
				slices.Delete(stack, len(stack)-1, len(stack))
			}
		} else if len(stack) == 0 {
			if _, ok := brackets[sVal]; !ok {
				return false
			}
		}

	}
	if len(stack) != 0 {
		return false
	}
	return true
}
