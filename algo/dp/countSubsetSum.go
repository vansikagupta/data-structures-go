/***** Explaination ****
Find number of subsets possible with a given sum
Variation of Subset Sum Problem
See subsetSum.go

****
To see if subset with given sum is possible we explored only till we got an answer.
But here we want to explore all possibilities.
Whenever we find a solution, increment count by 1
Video explaination: https://www.youtube.com/watch?v=MqYLmIzl8sQ&list=PLEJXowNB4kPxBwaXtRO1qFLpCzF75DYrS&index=10
*/

package dp

// Recursive simple solution

func countSubsetSum(items []int, sum, count int) int {
	if len(items) < 1 {
		if sum == 0 {
			return count + 1
		}
		return count
	}
	if items[0] > sum {
		return countSubsetSum(items[1:], sum, count)
	}
	// Exploring all posibilities
	return countSubsetSum(items[1:], sum, count) + countSubsetSum(items[1:], sum-items[0], count)
}

/* ** DP solution **
Store count in DP table
DP relation: lookup[i][j] = lookup[i-1][j] + lookup[i-1][j-items[0]]
*/
