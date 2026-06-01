package leetcoodee

// Question:
/*  1871. Jump Game VII
Medium

You are given a 0-indexed binary string s and two integers minJump and maxJump. In the beginning, you are standing at index 0, which is equal to '0'. You can move from index i to index j if the following conditions are fulfilled:

i + minJump <= j <= min(i + maxJump, s.length - 1), and
s[j] == '0'.
Return true if you can reach index s.length - 1 in s, or false otherwise.



Example 1:

Input: s = "011010", minJump = 2, maxJump = 3
Output: true
Explanation:
In the first step, move from index 0 to index 3.
In the second step, move from index 3 to index 5.
Example 2:

Input: s = "01101110", minJump = 2, maxJump = 3
Output: false


Constraints:

2 <= s.length <= 105
s[i] is either '0' or '1'.
s[0] == '0'
1 <= minJump <= maxJump < s.length**/

func CanReach(s string, minJump int, maxJump int) bool {
	if s[0] != '0' {

		return false
	}
	if s[len(s)-1] != '0' {
		return false
	}

	reachable := make([]bool, len(s))
	reachable[0] = true
	count := 0

	for j := 1; j < len(s); j++ {
		end := j - maxJump - 1
		start := j - minJump

		if end >= 0 && reachable[end] == true {
			count--
		}
		if start >= 0 && reachable[start] == true {
			count++
		}
		if count > 0 && s[j] == '0' {
			reachable[j] = true
		}
	}

	return reachable[len(reachable)-1]
}
