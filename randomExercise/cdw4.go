package main

import (
	"math"
	"strconv"
	"strings"
)

func AddByPower(s []string) int {
	indes := len(s) - 1
	total := 0
	for i := 0; i < len(s); i++ {
		num, _ := strconv.Atoi(s[i])
		total += num * int(math.Pow(float64(256), float64(indes)))
		indes--
	}
	return total
}

func Ip4(s, s1 string) int {
	newS := strings.Split(s, ".")
	newS1 := strings.Split(s1, ".")
	t1 := AddByPower(newS)
	t2 := AddByPower(newS1)
	if t1 >= t2 {
		return t1 - t2
	}
	return t2 - t1
}
