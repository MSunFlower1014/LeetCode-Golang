package main

import "fmt"

/*
9. 回文数
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
*/
func main() {
	result := isPalindrome(112211)
	fmt.Print(result)
}

func isPalindrome(x int) bool {
	//小于0或为10的倍数 都非回文数
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	rev := 0
	for x > rev {
		rev = rev*10 + x%10
		x = x / 10
	}
	return x == rev || x == rev/10
}
