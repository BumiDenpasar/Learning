package main

import "fmt"

func twoSum(nums []int, target int) []int {
	if len(nums) <= 1 {
		return []int{0, 1}
	}

	for i := 0; i < len(nums)-1 && nums[i] <= target; i++ {
		for n := i + 1; n < len(nums) && nums[n] <= target; n++ {
			if nums[n]+nums[i] == target {
				return []int{i, n}
			}
			fmt.Println("n=", n)
		}
		fmt.Println("i=", i)
	}
	return []int{0}
}

func main() {
	var nums = []int{0, 3, 4, 0}
	target := 0
	fmt.Println(twoSum(nums, target))
}
