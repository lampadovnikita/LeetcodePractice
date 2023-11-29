// https://leetcode.com/problems/backspace-string-compare/
// 844. Backspace String Compare
//
// Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.
// Note that after backspacing an empty text, the text will continue empty.
//
// Example 1:
// Input: s = "ab#c", t = "ad#c"
// Output: true
// Explanation: Both s and t become "ac".
//
// Example 2:
// Input: s = "ab##", t = "c#d#"
// Output: true
// Explanation: Both s and t become "".
//
// Example 3:
// Input: s = "a#c", t = "b"
// Output: false
// Explanation: s becomes "c" while t becomes "b".

package easy

func backspaceCompare(s string, t string) bool {
	sr := []rune(s)
	tr := []rune(t)

	sIdx := len(sr) - 1
	tIdx := len(tr) - 1

	var back int

	for {
		back = 0
		for (sIdx >= 0) && (back > 0 || sr[sIdx] == '#') {
			if sr[sIdx] == '#' {
				back++
			} else {
				back--
			}

			sIdx--
		}

		back = 0
		for (tIdx >= 0) && (back > 0 || tr[tIdx] == '#') {
			if tr[tIdx] == '#' {
				back++
			} else {
				back--
			}

			tIdx--
		}

		if (sIdx >= 0) && (tIdx >= 0) && (sr[sIdx] == tr[tIdx]) {
			sIdx--
			tIdx--
		} else {
			break
		}
	}

	if (sIdx < 0) && (tIdx < 0) {
		return true
	} else {
		return false
	}
}
