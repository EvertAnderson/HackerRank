package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		//fmt.Println(nums[i])
		tempTarget := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if tempTarget == nums[j] {
				fmt.Println(i, " ", j)
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

func main() {
	// x := [6]int{1, 2, 3, 4, 5, 8}
	// twoSum(x[:], 5)
	twoSum([]int{1, 2, 3, 4, 5, 8}, 5)
	fmt.Println("Hola Mundo")
}
