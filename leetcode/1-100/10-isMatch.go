package main

/*
10. 正则表达式匹配
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。

字节码
.   --->  46
*   --->  42
*/
func main() {
	result := isMatch("aabc", "a*b***c")
	print(result)
}

func isMatch(s string, p string) bool {
	index := 0
	flag := false
	for i, v := range p {
		if v == 46 {
			if i == len(p)-1 {
				index++
				break
			}
			index++
			continue
		}
		if v == 42 {
			if i == len(p)-1 {
				return true
			}
			flag = true
			continue
		}
		if index > len(s) {
			return false
		}
		if uint8(v) == s[index] {
			index++
			flag = false
			continue
		}
		if flag {
			if uint8(v) != s[index] {
				index++
				for true {
					if uint8(v) == s[index] {
						index++
						break
					}
					index++

				}
				flag = false
				continue
			} else {
				index++
				flag = false
				continue
			}
		}
		break
	}
	return index == len(s)
}
