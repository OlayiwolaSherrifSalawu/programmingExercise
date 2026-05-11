package main

import (
	"fmt"
	"strconv"
)

func ReadAbleTime(times int) string {
	hour := "00"
	min := "00"
	sec := "00"
	h := times / 60
	s := times % 60
	timess := ""
	if h == 0 {
		if s > 10 {
			sec = strconv.Itoa(s)
		} else if s <= 10 {
			sec = "0" + strconv.Itoa(s)
		}
		timess := fmt.Sprintf("%s:%s:%s", hour, min, sec)

		return timess
	}
	if h > 59 {
		st := h / 60
		sts := h % 60
		if st > 10 {
			hour = strconv.Itoa(st)
		} else if st < 10 {
			hour = "0" + strconv.Itoa(st)
		}
		if sts > 10 {
			min = strconv.Itoa(sts)
		} else if sts < 10 {
			min = "0" + strconv.Itoa(sts)
		}
		if s > 10 {
			sec = strconv.Itoa(s)
		} else if s < 10 {
			sec = "0" + strconv.Itoa(s)
		}
	}
	if h >= 1 && h <= 59 {
		if h > 10 {
			min = strconv.Itoa(h)
		} else if h < 10 {
			min = "0" + strconv.Itoa(h)
		}

		if s > 10 {
			sec = strconv.Itoa(s)
		} else if s < 10 {
			sec = "0" + strconv.Itoa(s)
		}
	}
	timess = fmt.Sprintf("%s:%s:%s", hour, min, sec)
	return timess
}
