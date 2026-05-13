package main

import "fmt"

func main() {
	val := nestedSum([]any{[]any{}, []any{1, []any{}}, 2})
	fmt.Println(val)
}
func nestedSum(items []any) int {
	theIndex := 0
	value := []any{}
	count := 0
	// base case for the the recursive function
	if len(items) == 1 {
		count = items[0].(int)
		return count
	}
	// to count int values alone
	for _, val := range items {
		if v, ok := val.(int); ok {
			count += v
		}
	}
	// get the first any[] slice to be able to put other value in it
	for i, val := range items {
		if v, ok := val.([]any); ok {
			value = v
			theIndex = i
			break
		}
	}

	value = append(value, count)
	for i, val := range items {
		if v, ok := val.([]any); ok && i != theIndex {
			value = append(value, v)
		}
	}
	items = value

	return nestedSum(items)
}
