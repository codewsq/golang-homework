package main

import "fmt"

/*
最长公共前缀

考察：字符串处理、循环嵌套

题目：查找字符串数组中的最长公共前缀
*/

func longestCommonPrefixVertical(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 逐个字符比较
	for i := 0; ; i++ {
		// 检查第一个字符串的第i个字符
		if i >= len(strs[0]) {
			return strs[0][:i] // [:i] 是切片操作，意思是取从开始到索引i（不包括i）的部分
		}
		// 获取第一个元素字符串 中 第 i 个字符
		currentChar := strs[0][i]

		// 检查其他字符串的第i个字符是否相同
		for j := 1; j < len(strs); j++ {
			// 如果当前字符串长度不够，或者字符不匹配
			if i >= len(strs[j]) || strs[j][i] != currentChar {
				return strs[0][:i]
			}
		}
	}
}
func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefixVertical(strs))

	strs = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefixVertical(strs))
}
