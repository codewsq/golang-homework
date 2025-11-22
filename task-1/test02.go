package main

import "fmt"

/*
有效的括号

考察：字符串处理、栈的使用

题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
*/

func isValid(s string) bool {
	// 如果字符串长度为奇数，肯定不匹配
	if len(s)%2 != 0 {
		return false
	}

	// 声明切片slice 假设是一个栈
	stack := make([]rune, 0)
	// 定义括号映射关系
	maps := make(map[rune]rune, 3)
	maps[')'] = '('
	maps[']'] = '['
	maps['}'] = '{'

	// 遍历字符串
	for _, char := range s {
		// 如果是右括号
		if _, ok := maps[char]; ok {
			// 如果栈中没有数据，或者 其数据不等于map中对应的左括号
			if len(stack) == 0 || stack[len(stack)-1] != maps[char] {
				return false
			} else {
				// 弹出栈顶元素(去掉最后一个元素)
				stack = stack[:len(stack)-1]
			}
		} else {
			// 将左括号保存在栈中
			stack = append(stack, char)
		}

	}
	// 最后栈应该为空，否则说明有未匹配的左括号
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("{[]}"))
}
