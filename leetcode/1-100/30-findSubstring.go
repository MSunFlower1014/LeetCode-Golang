package main

import "fmt"

/*
30. 串联所有单词的子串
给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。



示例 1：

输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：

输入：
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
输出：[]

"barfoofoobarthefoobarman"
["bar","foo","the"]
*/
func main() {
	result := findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"})
	fmt.Print(result)
}

func findSubstring(s string, words []string) []int {
	result := []int{}
	if len(words) == 0 {
		return result
	}
	length := len(words[0])
	subLength := length * len(words)

	if len(s) < subLength || len(s) == 0 || length == 0 {
		return result
	}
	index := 0
	var dict = make(map[string]int)
	for _, v := range words {
		dict[v]++
	}
	for i := range s {
		if len(s)-i < subLength {
			break
		}
		if checkWord(s, i, i+length, dict, index, subLength, length) {
			result = append(result, i)
		}
		index++
	}
	return result
}

func checkWord(s string, start, end int, dict map[string]int, index int, subLength int, length int) bool {
	temp := s[start:end]
	if dict[temp] > 0 {
		if end-index == subLength {
			return true
		}
		dict[temp]--
		if !checkWord(s, start+length, end+length, dict, index, subLength, length) {
			dict[temp]++
			return false
		}
		dict[temp]++
		return true
	} else {
		return false
	}
}
