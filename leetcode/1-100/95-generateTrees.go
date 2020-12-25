package main

/*
95. 不同的二叉搜索树 II
给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。



示例：

输入：3
输出：
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
解释：
以上的输出对应以下 5 种不同结构的二叉搜索树：

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/
func main() {

}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return treeHelp(1, n)
}

func treeHelp(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	var allTrees []*TreeNode
	for i := start; i <= end; i++ {
		//获取所有左子树
		leftTree := treeHelp(start, i-1)
		//获取所有右子树
		rightTree := treeHelp(i+1, end)

		for _, left := range leftTree {
			for _, right := range rightTree {
				currentTree := &TreeNode{i, nil, nil}
				currentTree.Left = left
				currentTree.Right = right
				allTrees = append(allTrees, currentTree)
			}
		}
	}
	return allTrees
}
