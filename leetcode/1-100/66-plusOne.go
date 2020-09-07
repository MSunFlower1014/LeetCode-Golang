package main

/*
66. 加一
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。
*/
func main() {

}

func plusOne(digits []int) []int {
	length := len(digits)
	temp := 0
	var result []int
	digits[length-1] = digits[length-1] + 1
	if digits[length-1] > 9 {
		temp = 1
		digits[length-1] = digits[length-1] % 10
	}
	for i := length - 2; i >= 0; i-- {
		if temp == 0 {
			break
		}
		digits[i] = digits[i] + temp
		if digits[i] > 9 {
			digits[i] = digits[i] % 10
			temp = 1
		} else {
			temp = 0
		}
	}

	if temp == 1 {
		result = append(result, 1)
		for _, v := range digits {
			result = append(result, v)
		}
	} else {
		result = digits
	}
	return digits
}
