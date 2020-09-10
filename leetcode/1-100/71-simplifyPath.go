package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(simplifyPath("/a/./b/../../c/"))
}

func simplifyPath(path string) string {
	var result []string
	for _, v := range strings.Split(path, "/") {
		switch v {
		case ".":
			break
		case "":
			break
		case "..":
			if l := len(result); l > 0 {
				result = result[:l-1]
			}
			break
		default:
			result = append(result, v)
		}
	}
	return "/" + strings.Join(result, "/")
}
