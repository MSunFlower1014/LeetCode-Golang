package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*

811. 子域名访问计数
一个网站域名，如"discuss.leetcode.com"，包含了多个子域名。作为顶级域名，常用的有"com"，下一级则有"leetcode.com"，最低的一级为"discuss.leetcode.com"。当我们访问域名"discuss.leetcode.com"时，也同时访问了其父域名"leetcode.com"以及顶级域名 "com"。

给定一个带访问次数和域名的组合，要求分别计算每个域名被访问的次数。其格式为访问次数+空格+地址，例如："9001 discuss.leetcode.com"。

接下来会给出一组访问次数和域名组合的列表cpDomains 。要求解析出所有域名的访问次数，输出格式和输入格式相同，不限定先后顺序。

*/
func main() {
	var domains []string = []string{"9001 discuss.leetcode.com", "50 yahoo.com"}
	result := subdomainVisits(domains)
	fmt.Print(result)
}

func subdomainVisits(cpDomains []string) []string {
	m := make(map[string]int)
	for _, row := range cpDomains {
		list := strings.Split(row, " ")
		time, _ := strconv.Atoi(list[0])
		domain := list[1]
		domainList := strings.Split(domain, ".")
		rowDomain := ""

		for i := len(domainList) - 1; i >= 0; i-- {
			if rowDomain == "" {
				rowDomain = domainList[i] + rowDomain
			} else {
				rowDomain = domainList[i] + "." + rowDomain
			}
			m[rowDomain] = m[rowDomain] + time
		}
	}
	var result []string
	for k, v := range m {
		lStr := strconv.Itoa(v)
		resultRow := lStr + " " + k
		result = append(result, resultRow)
	}
	return result
}
