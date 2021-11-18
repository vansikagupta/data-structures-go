// leetcode #1 https://leetcode.com/problems/two-sum/
// @hash-map @two-pointer

//Given an array of integers nums and an integer target,
//return indices of the two numbers such that they add up to target.

package hashmap

/*
* Complexity O(n2)
* traverse through the array and for each element,
* search for it's complement rest of the array
 */

func twoSum(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		//find the complement of nums[i]
		//linear search O(n) each time
		for j := 0; j < l; j++ {
			if i != j && (nums[j] == target-nums[i]) {
				return []int{i, j}
			}
		}
	}
	return nil
}

/*
* Improve the above algorithm by using binary search if the array is sorted
* search time reduces to log(n)
* Complexity O(nlogn)
 */

/* Two-pointer
* For a sorted array, start from the first and last indices
* if sum > target, move left at the end (excluding the current largest element)
* if sum < target, move right at the front (excluding the current smallest element)
* if sum == target, done
 */

// assuming we have sorted the array in O(nlogn) time
// Complexity O(nlogn)

func twoSumTwoPointer(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	//excluding the possibility of both indices pointing to the same element
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			// both compliment each other very well
			return []int{left + 1, right + 1}
		}
		if sum > target {
			// we have to get rid of the current largest element
			// because it is complaining with the smallest possible compliment also
			// cannot have any possible compliment
			right = right - 1
		} else {
			left = left + 1
		}
	}
	return nil
}

/*
* Use hash table to search for compelment of a number
* Extra space required
* Complexity time O(n), space O(n)
 */

func twoSumHash(nums []int, target int) []int {
	hash := make(map[int]int)
	// first parse to insert all numbers in map
	for i, v := range nums {
		hash[v] = i
	}

	//second parse; for num i, search it's complement, until found
	for i, v := range nums {
		comp := target - v
		// if complement of nums[i] is present, return the indices
		if index, ok := hash[comp]; ok && i != index {
			return []int{i, index}
		}
	}
	return nil
}

// This can be done in one pass if order of index doesn't matter.
