package main

import (
	"fmt"
	"strconv"
)

/*
输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
加 1 后得到 4321 + 1 = 4322。
因此，结果应该是 [4,3,2,2]。

4329 + 1 = 4330 [4,3,3,0]
*/
func numberArray(numbers []int) []int {
	if len(numbers) == 0 {
		return numbers
	}

	// 获取数组中的最后一位数据
	number := numbers[len(numbers)-1]
	if number == 9 { // 如果最后一位为9
		sumStr := ""
		// 将数组元素按顺序拼接为字符串
		for _, value := range numbers {
			sumStr = sumStr + strconv.Itoa(value)
		}
		// 转换字符串为整数并加1
		sum, _ := strconv.Atoi(sumStr)
		sum = sum + 1
		// 将计算结果转换为字符串结果
		sumStr = strconv.Itoa(sum)
		// 重新初始化数组长度
		numbers = make([]int, len(sumStr))
		// 将字符串结果逐次转换为数组元素
		for index, value := range sumStr {
			num, _ := strconv.Atoi(string(value))
			numbers[index] = num
		}
	} else {
		numbers[len(numbers)-1] = numbers[len(numbers)-1] + 1
	}

	return numbers
}

func main() {
	numbers := []int{9}
	fmt.Println(numberArray(numbers))

	numbers = []int{4, 3, 2, 1}
	fmt.Println(numberArray(numbers))

	numbers = []int{4, 3, 2, 9}
	fmt.Println(numberArray(numbers))

	numbers = []int{4, 9, 9, 9}
	fmt.Println(numberArray(numbers))

	numbers = []int{9, 9, 9, 9}
	fmt.Println(numberArray(numbers))
}
