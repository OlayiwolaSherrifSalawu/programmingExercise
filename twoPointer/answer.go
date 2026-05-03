package twopointer

// practising the tortise and hear algorithm
type WayPoint struct {
	Data string
	next *WayPoint
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

func DetectCollision(head *WayPoint) bool {
	tortise := head
	hear := head
	for hear.next != nil || hear != nil {
		hear = hear.next.next
		tortise = tortise.next

		if tortise == hear {
			return true
		}
	}
	return false
}
