package main

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

func nestedSums(items []any) int {
	total := 0

	for _, val := range items {
		if s, ok := val.(int); ok {
			total += s
		}
		if s, ok := val.([]any); ok {
			total += nestedSum(s)
		}
	}
	return total
}
func nestedSum1(items []any) int {
	
}
