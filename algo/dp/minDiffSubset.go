// interviewbit Q:https://www.interviewbit.com/problems/minimum-difference-subsets/
/*** Problem ***
* Given a set of positive integers, find the minimum difference two subsets, if partitioned, can have

Solution:
If equal partition is possible, s1 = s2 = sum/2 and diff = 0
Otherwise, one subset will be greater and the other will be smaller, such that
s1 < sum/2 < s2 and s1 + s2 = sum

If we are able to find one of the sets, the other one can be easily deduced from the main set
s1 = sum - s2 and s2 = sum - s1

So let's now reduce the problem to find just one such subset such that s1 < sum/2
Rephrasing the problem now
Find subset from a given set such that sum of subset <= sum/2
Sounds like 0/1 knapsack :)
*/

package dp

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// We are trying to find the maximu sum a subset can have such that the max sum < given sum
func maxSumSubset(items []int, sum int) int {
	if len(items) < 1 || sum == 0 {
		return 0
	}
	// This check right here is making sure that capacity cannot be exceeded
	if items[0] > sum {
		return maxSumSubset(items[1:], sum)
	}
	return max(maxSumSubset(items[1:], sum), items[0]+maxSumSubset(items[1:], sum-items[0]))
}

func maxSumSubsetTab(items []int, sum int) int {
	n := len(items)
	lookup := make([][]int, n+1)
	for i := range lookup {
		lookup[i] = make([]int, sum+1)
	}
	for k := 0; k <= sum; k++ {
		lookup[0][k] = 0
	}
	for k := 0; k <= n; k++ {
		lookup[k][0] = 0
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if items[i-1] > j {
				lookup[i][j] = lookup[i-1][j]
			} else {
				lookup[i][j] = max(lookup[i-1][j], items[i-1]+lookup[i-1][j-items[i-1]])
			}
		}
	}
	// lookup[n][sum] is going to hold the max sum
	return lookup[n][sum]
}

func minDiffSubset(set []int) int {
	sum := 0
	for _, val := range set {
		sum = sum + val
	}
	s1 := maxSumSubset(set, sum/2)
	return sum - 2*s1
}
