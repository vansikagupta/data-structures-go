// leetcode #416: https://leetcode.com/problems/partition-equal-subset-sum/

/***** Explaination ****
Variation of Subset Sum Problem
See subsetSum.go

****
In order to partition given set into two equal sum subsets,
the given set must have sum as multiple of 2, say it is 2x.

Now we need to figure out if a subset with sum as x exists in the given set.
If we are able to find solution for one subset, it implies that the other subset will also have sum as x (as remaining sum is 2x -x)
*/
package dp

import "math"

func canPartition(nums []int) bool {
	setSum := 0
	for _, val := range nums {
		setSum += val
	}
	//Set can be partitioned into two eual sum subsets only if sum of the set is multiple of 2
	if math.Mod(float64(setSum), 2) > 0.0 {
		return false
	}
	// Now find if given set contains a subset with sum = setSum/2
	// Subset sum problem
	return subsetSumTab(nums, setSum/2)
}
