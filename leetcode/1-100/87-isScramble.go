package main

/*
87. 扰乱字符串
给定一个字符串 s1，我们可以把它递归地分割成两个非空子字符串，从而将其表示为二叉树。

下图是字符串 s1 = "great" 的一种可能的表示形式。

    great
   /    \
  gr    eat
 / \    /  \
g   r  e   at
           / \
          a   t
在扰乱这个字符串的过程中，我们可以挑选任何一个非叶节点，然后交换它的两个子节点。

例如，如果我们挑选非叶节点 "gr" ，交换它的两个子节点，将会产生扰乱字符串 "rgeat" 。

    rgeat
   /    \
  rg    eat
 / \    /  \
r   g  e   at
           / \
          a   t
我们将 "rgeat” 称作 "great" 的一个扰乱字符串。

同样地，如果我们继续交换节点 "eat" 和 "at" 的子节点，将会产生另一个新的扰乱字符串 "rgtae" 。

    rgtae
   /    \
  rg    tae
 / \    /  \
r   g  ta  e
       / \
      t   a
我们将 "rgtae” 称作 "great" 的一个扰乱字符串。

给出两个长度相等的字符串 s1 和 s2，判断 s2 是否是 s1 的扰乱字符串。

示例 1:

输入: s1 = "great", s2 = "rgeat"
输出: true
*/
func main() {

}

func isScramble(s1 string, s2 string) bool {
	n := len(s1)
	m := len(s2)
	if n != m {
		return false
	}
	/*
		dp[i][j][k][h] 表示 T[k..h]T[k..h] 是否由 S[i..j]S[i..j] 变来。
		由于变换必须长度是一样的，因此这边有个关系 j - i = h - kj−i=h−k ，
		可以把四维数组降成三维。dp[i][j][len]
		dp[i][j][len]
		表示从字符串 SS 中 ii 开始长度为 lenlen 的字符串是否能变换为从字符串 TT 中 jj 开始长度为 lenlen 的字符串

		作者：jerry_nju
		链接：https://leetcode-cn.com/problems/scramble-string/solution/miao-dong-de-qu-jian-xing-dpsi-lu-by-sha-yu-la-jia/
		来源：力扣（LeetCode）
		著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
	*/
	dp := make([][][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]bool, n+1)
			if s1[i] == s2[j] {
				dp[i][j][1] = true
			}
		}
	}

	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			for j := 0; j <= n-l; j++ {
				for k := 1; k <= l-1; k++ {
					if dp[i][j][k] && dp[i+k][j+k][l-k] {
						dp[i][j][l] = true
						break
					}

					if dp[i][j+l-k][k] && dp[i+k][j][l-k] {
						dp[i][j][l] = true
						break
					}
				}
			}
		}
	}
	return dp[0][0][n]
}
