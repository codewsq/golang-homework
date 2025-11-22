package main

import "fmt"

/*
136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
	可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素

考察：数字操作、条件判断
题目：判断一个整数是否是回文数
*/

func main() {
	// 给定一个非空整数数组
	numbers := []string{"a", "a", "b", "b", "c", "c", "d", "e", "e"}

	numberMap := make(map[string]int)

	for i := 0; i < len(numbers); i++ {
		// 判断map中是否存在该元素，若存在，则加一
		if _, ok := numberMap[numbers[i]]; ok {
			numberMap[numbers[i]] = numberMap[numbers[i]] + 1
		} else {
			numberMap[numbers[i]] = 1
		}
	}

	for k, v := range numberMap {
		if v == 1 {
			fmt.Println("只出现一次的元素为：", k)
		}
	}
}
