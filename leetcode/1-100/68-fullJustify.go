package main

import "strings"

/*
68. 文本左右对齐
给定一个单词数组和一个长度 maxWidth，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。

你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。

要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。

文本的最后一行应为左对齐，且单词之间不插入额外的空格。

说明:

单词是指由非空格字符组成的字符序列。
每个单词的长度大于 0，小于等于 maxWidth。
输入单词数组 words 至少包含一个单词。
示例:

输入:
words = ["This", "is", "an", "example", "of", "text", "justification."]
maxWidth = 16
输出:
[
   "This    is    an",
   "example  of text",
   "justification.  "
]
*/
func main() {

}

/*
作者：Triste_24
链接：https://leetcode-cn.com/problems/text-justification/solution/xing-kuan-du-chao-chu-xing-zhi-you-yi-ge-dan-ci-zu/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func fullJustify(words []string, maxWidth int) []string {
	//make一个空行后面使用
	var strBuilder strings.Builder
	for i := 0; i < maxWidth; i++ {
		strBuilder.WriteString(" ")
	}
	emptyLine := strBuilder.String()

	ret := make([]string, 0)
	count, start := 0, 0
	for i := 0; i < len(words); i++ {
		word := words[i]
		count += len(word)
		if count > maxWidth {
			//宽度超出，处理[start,i-1]的单词
			ret = append(ret, makeLine(words, emptyLine, start, i-1))
			start = i
			if i == len(words)-1 {
				//如果此时正好还是最后一个单词，重新处理最后一个单词
				count = 0
				i--
			} else {
				count = len(word) + 1
			}
		} else if i == len(words)-1 {
			//最后一个单词，且没超限，处理[start,i]的单词
			ret = append(ret, makeLine(words, emptyLine, start, i))
		} else {
			//添上一个空格
			count++
		}
	}
	return ret
}

func makeLine(words []string, emptyLine string, start, end int) string {
	chars := []byte(emptyLine)
	offset := 0
	//一个单词或者最后一行，左对齐
	if start == end || end == len(words)-1 {
		for i := start; i <= end; i++ {
			copy(chars[offset:], words[i])
			//只需留一个空格
			offset += len(words[i]) + 1
		}
		return string(chars)
	}

	wordCount := end - start + 1
	left := len(emptyLine) - wordCount + 1
	for i := start; i <= end; i++ {
		left -= len(words[i])
	}
	spaceCount := left / (wordCount - 1)
	mod := left % (wordCount - 1)
	for i := start; i <= end; i++ {
		copy(chars[offset:], words[i])
		//留下1+商的空格
		offset += len(words[i]) + 1 + spaceCount
		if mod > 0 {
			//再多留一个模的的空格
			offset++
			mod--
		}
	}
	return string(chars)
}
