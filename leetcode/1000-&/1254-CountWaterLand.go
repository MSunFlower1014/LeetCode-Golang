package main

/*
1254. 统计封闭岛屿的数目
有一个二维矩阵 grid ，每个位置要么是陆地（记号为 0 ）要么是水域（记号为 1 ）。

我们从一块陆地出发，每次可以往上下左右 4 个方向相邻区域走，能走到的所有陆地区域，我们将其称为一座「岛屿」。

如果一座岛屿 完全 由水域包围，即陆地边缘上下左右所有相邻区域都是水域，那么我们将其称为 「封闭岛屿」。

请返回封闭岛屿的数目。
{1,1,1,1,1,1,1,0}
{1,0,0,0,0,1,1,0}
{1,0,1,0,1,1,1,0}
{1,0,0,0,0,1,0,1}
{1,1,1,1,1,1,1,0}

1 1 1 1
1 1 1 1
1 1 1 1
1 1 1 1
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-closed-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	grid := [][]int{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}}
	i := closedIsland(&grid)
	println(i)
}

func closedIsland(grid *[][]int) int {
	xen := len(*grid)
	yen := len((*grid)[0])
	i := 0
	for x := 0; x < xen; x++ {
		for y := 0; y < yen; y++ {
			if (*grid)[x][y] == 0 && dfs(grid, x, y) {
				i++
			}
		}
	}
	return i
}

func dfs(grid *[][]int, x int, y int) (b bool) {
	xen := len(*grid)
	yen := len((*grid)[0])
	if x < 0 || y < 0 || x >= xen || y >= yen {
		return false
	}

	if (*grid)[x][y] == 1 {
		return true
	}
	(*grid)[x][y] = 1
	up := dfs(grid, x, y-1)
	down := dfs(grid, x, y+1)
	left := dfs(grid, x-1, y)
	right := dfs(grid, x+1, y)
	if up && down && left && right {
		return true
	}
	return false
}
