// https://leetcode.com/problems/climbing-stairs/
// 70. Climbing Stairs
//
// You are climbing a staircase. It takes n steps to reach the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
// Example 1:
// Input: n = 2
// Output: 2
// Explanation: There are two ways to climb to the top.
// 1. 1 step + 1 step
// 2. 2 steps
//
// Example 2:
// Input: n = 3
// Output: 3
// Explanation: There are three ways to climb to the top.
// 1. 1 step + 1 step + 1 step
// 2. 1 step + 2 steps
// 3. 2 steps + 1 step

package easy

// n1 - the number of ways to climb 2 steps above
// n2 - the number of ways to climb 1 step above
// We can climb the stair from [n-1] = n2 stair
// or from [n-2] = n1 stair
// From [n-1] there is one possible step and
// from [n-2] two different ways.
// So the total amount of ways = n1 + n2
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	n1 := 1
	n2 := 2

	for i := 3; i <= n; i++ {
		tmp := n2

		n2 += n1

		n1 = tmp
	}

	return n2
}
