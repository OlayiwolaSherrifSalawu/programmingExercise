package main

import "fmt"

func main() {

	fmt.Println(removeZeros([]int{0, 0, 1, 2}))

	fmt.Println(removeZeros([]int{1, 2, 0, 3, 0}))
	fmt.Println(removeZeros([]int{0}))
	fmt.Println(removeZeros([]int{1, 2, 3}))

}
