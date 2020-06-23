package main

import "fmt"

/*
20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
*/
func main() {
	print("")
	fmt.Print(isValid("()"))
}
func isValid(s string) bool {
	m := map[uint8]uint8{'(': ')', '[': ']', '{': '}'}
	var queue []uint8

	for i := range s {
		length := len(queue)
		if i == 0 || length == 0 {
			queue = append(queue, s[i])
			continue
		}
		if s[i] == m[queue[length-1]] {
			queue = queue[0 : length-1]
		} else {
			queue = append(queue, s[i])
		}
	}
	if len(queue) > 0 {
		return false
	} else {
		return true
	}
}
