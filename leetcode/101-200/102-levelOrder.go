package main

/*
102. 二叉树的层序遍历
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。



示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层序遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
*/
func main() {

}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var levelNode []*TreeNode
	levelNode = append(levelNode, root)
	for true {
		var tempLevelNode []*TreeNode
		var tempLevel []int
		for _, v := range levelNode {
			tempLevel = append(tempLevel, v.Val)
			if v.Left != nil {
				tempLevelNode = append(tempLevelNode, v.Left)
			}
			if v.Right != nil {
				tempLevelNode = append(tempLevelNode, v.Right)
			}
		}
		if len(tempLevel) > 0 {
			result = append(result, tempLevel)
			levelNode = tempLevelNode
		} else {
			break
		}
	}
	return result
}
