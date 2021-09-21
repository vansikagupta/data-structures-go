// leetcode #416: https://leetcode.com/problems/partition-equal-subset-sum/

package dp

import "math"

/***** Explaination ****

 */

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
