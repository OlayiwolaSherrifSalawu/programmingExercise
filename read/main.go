package main

import (
	"strconv"
	"strings"
)

func main() {
	// longest sub string in a string

}
func rpnCalc(s string) string{
	operator:= "-+/%*"
	rpn := strings.Fields(s)
	stack := []int{}
	for i:= 0; i < len(rpn); i++{
		op:= ""
		nu, err:= strconv.Atoi(rpn[i])
		if err != nil{
			if strings.ContainsAny(operator, rpn[i]){

			}
		}
	}

}