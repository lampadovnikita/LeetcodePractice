// https://leetcode.com/problems/is-subsequence/
// 392. Is Subsequence
//
// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
// A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
//
// Example 1:
// Input: s = "abc", t = "ahbgdc"
// Output: true
//
// Example 2:
// Input: s = "axc", t = "ahbgdc"
// Output: false

package easy

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}

	sr := []rune(s)
	tr := []rune(t)

	sIdx := 0
	tIdx := 0

	for tIdx < len(tr) {
		if sr[sIdx] == tr[tIdx] {
			sIdx++

			if sIdx >= len(sr) {
				return true
			}
		}

		tIdx++
	}

	return false
}
