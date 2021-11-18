package hashmap

/*
Two traversals
*/
func solve(A []int) int {
	hashmap := make(map[int]int)

	for _, v := range A {
		if _, ok := hashmap[v]; ok {
			hashmap[v] = hashmap[v] + 1
		} else {
			hashmap[v] = 1
		}

	}

	for _, v := range A {
		if hashmap[v] > 1 {
			return v
		}
	}

	return -1

}

/*
One tarversal
Maintain min-index that will hold the index of the first duplicate
*/

func solveOnePass(nums []int) int {
	minIndex := len(nums) // invalid index for sure
	hashmap := make(map[int]int, minIndex)

	for i, v := range nums {
		if index, ok := hashmap[v]; ok {
			minIndex = min(index, minIndex)
		} else {
			hashmap[v] = i
		}
	}

	if minIndex < len(nums) {
		return nums[minIndex]
	}
	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
func main() {
	fmt.Println(solve([]int{10, 5, 3, 5, 6}))
}*/
