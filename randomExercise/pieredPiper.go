package main

func PieredPiper(s string) int {
	head := byte('O')
	tail := byte('~')
	foundPiper := false
	astray := 0

	for i := 0; i < len(s); i += 2 {
		if rune(s[i]) == ' ' {
			i--
			continue
		}
		if rune(s[i]) == 'P' {
			foundPiper = true
			i--
			continue
		}
		if i+1 < len(s) {
			if s[i] == head && s[i+1] == tail && !foundPiper {
				astray++
			}
			if s[i] == tail && s[i+1] == head && foundPiper {
				astray++
			}
		}
	}
	return astray
}
