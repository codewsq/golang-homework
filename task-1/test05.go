package main

import "fmt"

/*
26. 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，
你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，
当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
*/

func handlerArrayNums(nums []int) int {
	// 如果数组为空或只有一个元素，直接返回长度
	if len(nums) <= 1 {
		return len(nums)
	}

	// 慢指针，指向当前不重复元素的最后一个位置
	i := 0

	// 快指针，遍历整个数组
	for j := 1; j < len(nums); j++ {
		// 如果当前元素与慢指针指向的元素不同
		if nums[j] != nums[i] {
			// 将慢指针后移
			i++
			// 将快指针指向的值赋给慢指针位置
			nums[i] = nums[j]
		}
		// 如果相同，快指针继续前进，慢指针不动
	}
	// 新数组的长度（慢指针位置 + 1）
	nums = nums[:i+1]
	fmt.Println("操作后数组元素：", nums)
	// 返回新数组的长度（慢指针位置 + 1）
	return len(nums)
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println("不重复元素的长度：", handlerArrayNums(nums))

	fmt.Println("----------------------------------------")

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println("不重复元素的长度：", handlerArrayNums(nums))
}
