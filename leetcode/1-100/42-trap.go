package main

import "fmt"

/*
42. 接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
*/
func main() {
	fmt.Print(trap2([]int{2, 0, 2}))
}

func trap(height []int) int {
	if len(height) < 2 {
		return 0
	}
	target := 0
	for i := 0; i < len(height); i++ {
		maxLeft, maxRight := 0, 0
		for j := i; j >= 0; j-- {
			if height[j] > maxLeft {
				maxLeft = height[j]
			}
		}
		for j := i; j < len(height); j++ {
			if height[j] > maxRight {
				maxRight = height[j]
			}
		}
		if maxRight > maxLeft {
			target = target + maxLeft - height[i]
		} else {
			target = target + maxRight - height[i]
		}
	}
	return target
}

/*
和方法 2 相比，我们不从左和从右分开计算，我们想办法一次完成遍历。
从动态编程方法的示意图中我们注意到，只要 \text{right\_max}[i]>\text{left\_max}[i]right_max[i]>left_max[i]
（元素 0 到元素 6），积水高度将由 left_max 决定，
类似地 \text{left\_max}[i]>\text{right\_max}[i]left_max[i]>right_max[i]（元素 8 到元素 11）。
所以我们可以认为如果一端有更高的条形块（例如右端），积水的高度依赖于当前方向的高度（从左到右）。
当我们发现另一侧（右侧）的条形块高度不是最高的，我们则开始从相反的方向遍历（从右到左）。
我们必须在遍历时维护 \text{left\_max}left_max 和 \text{right\_max}right_max ，
但是我们现在可以使用两个指针交替进行，实现 1 次遍历即可完成。

作者：LeetCode
链接：https://leetcode-cn.com/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func trap2(height []int) int {
	left, right := 0, len(height)-1
	ans := 0

	leftMax, rightMax := 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				ans = ans + leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				ans = ans + rightMax - height[right]
			}
			right--
		}
	}
	return ans
}
