package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
8. 字符串转换整数 (atoi)
请你来实现一个 atoi 函数，使其能将字符串转换成整数。

首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。接下来的转化规则如下：

如果第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字字符组合起来，形成一个有符号整数。
假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成一个整数。
该字符串在有效的整数部分之后也可能会存在多余的字符，那么这些字符可以被忽略，它们对函数不应该造成影响。
注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换，即无法进行有效转换。

在任何情况下，若函数不能进行有效的转换时，请返回 0 。

提示：

本题中的空白字符只包括空格字符 ' ' 。
假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。如果数值超过这个范围，请返回  INT_MAX (231 − 1) 或 INT_MIN (−231) 。
字节码对应
+  --> 43
- ---> 45
空格  --> 32
0-9 --> 48-57
*/
func main() {
	result := myAtoi("0-1")
	fmt.Print(result)
}

func myAtoi(str string) int {
	flag := ""
	result := ""
	no := true
	for _, v := range str {
		if no && v == 32 {
			continue
		}
		if no && v == 43 {
			no = false
			continue
		}
		if no && v == 45 {
			no = false
			flag = "-"
			continue
		}
		if v >= 48 && v <= 57 {
			no = false
			result = result + string(v)
			continue
		}
		break
	}
	if result == "" {
		return 0
	}
	i, _ := strconv.Atoi(flag + result)
	if i > math.MaxInt32 {
		i = math.MaxInt32
	}
	if i < math.MinInt32 {
		i = math.MinInt32
	}
	return i
}
