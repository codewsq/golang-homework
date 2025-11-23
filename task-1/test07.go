package main

import "fmt"

/*
两数之和
考察：数组遍历、map使用
题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
*/

func numAddToTarget(nums []int, target int) (resultMap map[int]int) {
	if len(nums) == 0 {
		return map[int]int{}
	}

	resultMap = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		num1 := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j]+num1 == target {
				resultMap[nums[i]] = nums[j]
			}
		}
	}
	return resultMap
}

func main() {
	nums := []int{1, 2, 3, 7, 5, 8}

	target := 10

	resultMap := numAddToTarget(nums, target)
	fmt.Println(resultMap)
}
