package main

import "fmt"

/*
17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。



示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
*/
func main() {
	fmt.Print(letterCombinations("234"))
}

func letterCombinations(digits string) []string {
	m := make(map[string][]string)
	m["2"] = []string{"a", "b", "c"}
	m["3"] = []string{"d", "e", "f"}
	m["4"] = []string{"g", "h", "i"}
	m["5"] = []string{"j", "k", "l"}
	m["6"] = []string{"m", "n", "o"}
	m["7"] = []string{"p", "q", "r", "s"}
	m["8"] = []string{"t", "u", "v"}
	m["9"] = []string{"w", "x", "y", "z"}
	var result []string
	for i := 0; i < len(digits); i++ {
		var number string
		if i == len(digits)-1 {
			number = digits[i:]
		} else {
			number = digits[i : i+1]
		}
		var temp []string
		for _, v := range m[number] {
			if i == 0 {
				result = append(result, v)
				continue
			}

			for _, t := range result {
				temp = append(temp, t+v)
			}
		}
		if temp != nil {
			result = temp
		}

	}
	return result
}
