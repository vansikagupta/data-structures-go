// leetcode #287: https://leetcode.com/problems/find-the-duplicate-number/
/*
************
* This is a classic problem with multiple approaches.
* Constaints:
* 	Length of array is n+1
* 	Value of elements is in the range 1 to n
*   There's exactly one element that has duplicate
* Follow up:
* Proof that there has to be atleast one duplicate
* PigeonHole principle proves that atleast one duplicate must exists
* This principle states that of there be n+1 pigeons and n pigeon holes,
* then atleast one pigeon hole must have two or more pigeons.
************
 */

package arrays

/* Approach 1
* Sort the array O(nlogn)
* Find duplicate in sorted array O(n)
Time: O(nlogn)
Space: Complexity of the sorting algorithm
*/

/* Approach 2
* Remember the visited elements with Set
* If element is already present in set, duplicate found
Time: O(n)
Space: O(n)
*/
func findDuplicateSet(nums []int) int {
	n := len(nums)
	seen := make(map[int]bool, n)

	for _, element := range nums {
		//if seen the element before, duplicate found

		if seen[element] {
			return element
		}
		//Otherwise, remember the newly visited element
		seen[element] = true
	}
	return -1
}

/* Approach 3
* Since all elements are in the range of array index, we can use the negative marking approach
* There will be two such element(duplicates) that will point to the same index
* We use this property to identify the duplicate
* At each element index, change the sign
* If we find the element at that index already negative, we have found our duplicate
Time: O(n)
Space: O(1)
*********
* Modifies elements in-place
*/
func findDuplicateNegativeMarking(nums []int) int {
	for _, element := range nums {
		//omitting sign so that it can be used as index
		if element < 0 {
			element *= -1
		}
		if nums[element] < 0 {
			return element
		}
		nums[element] *= -1
	}
	return -1
}

// Brute force
// for each element, see if it's duplicate is present
// O(n^2) time
// O(1)
// Time Limit Execeeded
func findDuplicate(nums []int) int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	// should not be reached stated that there's always atleast one duplicate
	return -1
}
