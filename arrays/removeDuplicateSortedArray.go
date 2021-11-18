// leetcode #26: https://leetcode.com/problems/remove-duplicates-from-sorted-array/

/*****
* Remove duplicates in-place
 */

package arrays

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 || n == 1 {
		return n
	} // no duplicates possible
	k := 0
	for i := 0; i <= n-2; i++ {
		if nums[i] != nums[i+1] {
			nums[k+1] = nums[i+1]
			k++
		}
	}

	return k + 1
}
