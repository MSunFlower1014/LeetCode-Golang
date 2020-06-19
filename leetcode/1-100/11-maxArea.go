package main

/*
11. 盛最多水的容器
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。
*/
func main() {
	result := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	print(result)
}

func maxArea(height []int) (max int) {
	start, end := 0, len(height)-1
	for true {
		if start == end {
			break
		}
		var c int
		if height[start] > height[end] {
			c = height[end]
			end--
		} else {
			c = height[start]
			start++
		}
		if max < (end-start+1)*c {
			max = (end - start + 1) * c
		}
	}
	return
}
