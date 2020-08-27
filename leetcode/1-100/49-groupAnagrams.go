package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	var result [][]string

	dict := make(map[string][]string)

	for _, v := range strs {

		var temp []int
		for _, l := range v {
			temp = append(temp, int(l))
		}
		sort.Ints(temp)
		k := ""
		for _, v := range temp {
			k = k + (string(v))
		}

		if dict[k] == nil {
			var list []string
			list = append(list, v)
			dict[k] = list
		} else {
			dict[k] = append(dict[k], v)
		}
	}

	for _, v := range dict {
		result = append(result, v)
	}
	return result
}
