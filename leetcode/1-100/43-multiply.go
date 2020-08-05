package main

import "fmt"

/*
43. 字符串相乘
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
*/
func main() {
	fmt.Print(multiply("123", "20"))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	// num1[i] * num[j] 值必定在resArr[i+j] resArr[i+j+1]上，i+j+1存储地位
	resArr := make([]int, len(num1)+len(num2))
	for i := len(num2) - 1; i >= 0; i-- {
		n2 := int(num2[i] - '0')
		for j := len(num1) - 1; j >= 0; j-- {
			n1 := int(num1[j] - '0')
			sum := n2*n1 + resArr[i+j+1]
			resArr[i+j+1] = sum % 10
			resArr[i+j] += sum / 10
		}
	}

	res := ""
	for k, v := range resArr {
		if k == 0 && v == 0 {
			continue
		}
		res += string(v + '0')
	}

	return res
}
