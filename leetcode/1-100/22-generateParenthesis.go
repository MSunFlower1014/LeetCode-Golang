package main

import "fmt"

func main() {
	fmt.Print(generateParenthesis(5))
}

var caches = make(map[int][]string)

func generateParenthesis(n int) []string {
	return generate(n)
}

func generate(n int) []string {
	if caches[n] != nil {
		return caches[n]
	}

	var ans []string
	if n == 0 {
		ans = append(ans, "")
	} else {
		for c := 0; c < n; c++ {
			for _, left := range generate(c) {
				for _, right := range generate(n - 1 - c) {
					ans = append(ans, "("+left+")"+right)
				}
			}
		}
	}
	caches[n] = ans
	return ans
}
