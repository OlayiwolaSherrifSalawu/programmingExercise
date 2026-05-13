package main

type Report struct {
	EvenIndexSum int
	ScoreTotal   int
	RuneCount    int
}

func rangeReport(nums []int, scores map[string]int, word string) Report {
	Reports := Report{}
	// 1. sum all even indexes even zero
	for i, val := range nums {
		if i == 0 || i%2 == 0 {
			Reports.EvenIndexSum += val
		}
	}
	for _, val := range scores {
		Reports.ScoreTotal += val
	}
	for range word {
		Reports.RuneCount += 1
	}
	return Reports
}
