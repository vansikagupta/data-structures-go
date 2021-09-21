// interviewbit Q:https://www.interviewbit.com/problems/subset-sum-problem/

package dp

/*
Classic Example of 0/1 knapsack
Subset is my knapsack and bag capacity is my subset sum.
Each item can be included once only.
States: remaining items and remaining sum; f(n, sum)
Base Condition:
	Sum becomes zero.
Decision Tree:
	1. If item > remaning sum; item cannot be included; solution = subproblem solution
		f(n, sum) = f(n-1, sum)
	2. Else you may or may not include item; whichever gives you solution
	    f(n-1, sum) || f(n-1, sum-item)
*/

/**
 * @input items : Integer array
 * @input sum  : Integer
 *
 * @Output Integer
 * 1 true
 * 0 false
 */

// Plain recursion solution
// Time Complexity: O(2^n)
// Additional Stack space used

func subsetSum(items []int, sum int) bool {
	// base condition; positive solution
	if sum == 0 {
		return true
	}
	// negative solution; sum is not zero yet but no more items left
	if len(items) < 1 {
		return false
	}
	// there's still some items and sum left;
	// decision tree
	// exclude item
	if items[0] > sum {
		// f(n-1, sum)
		return subsetSum(items[1:], sum)
	}
	// may or may not include; whichever gives the solution first
	// f(n-1, sum) || f(n-1, sum-item)
	return subsetSum(items[1:], sum) || subsetSum(items[1:], sum-items[0])
}

/*
******** Convert Recursive solution to Tabulation *******
* add lookup table
* fill in base values
* 1. Since we have 2 states, our lookup table would be a 2D array
* 2. Fill lookup table with base cases:
	 when sum = 0, result is true
	 when num of items = 0, sum cannot be build, result is false
* 3. Use smaller solutions to solve larger problems
* 4. Decision Tree Relation:
	a. If item[i] > j (remaining sum); item cannot be included; solution = subproblem solution
		lookup[i, j] = lookup[i-1, j]
	b. Else you may or may not include item; whichever gives you solution
	    lookup[i, j] = lookup[i-1, j] || lookup[i-1, j-item[i]]


// Time Complexity: O(n*sum)
// Space: O(n*sum)
*/

func subsetSumTab(items []int, sum int) bool {
	// Step 1
	lookup := make([][]bool, len(items)+1)
	for i := range lookup {
		lookup[i] = make([]bool, sum+1)
	}
	// Step 2
	for k := 0; k <= sum; k++ {
		lookup[0][k] = false
	}
	for k := 0; k <= len(items); k++ {
		lookup[k][0] = true
	}
	// Step 3
	for i := 1; i <= len(items); i++ {
		for j := 1; j <= sum; j++ {
			if items[i-1] > j {
				lookup[i][j] = lookup[i-1][j] // exclude item
			} else {
				lookup[i][j] = lookup[i-1][j] || lookup[i-1][j-items[i-1]]
			}
		}
	}
	return lookup[len(items)][sum]
}
