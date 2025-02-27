# [783. Minimum Distance Between BST Nodes](https://leetcode.com/problems/minimum-distance-between-bst-nodes/)


## 题目

Given the `root` of a Binary Search Tree (BST), return *the minimum difference between the values of any two different nodes in the tree*.

**Note:** This question is the same as 530: [https://leetcode.com/problems/minimum-absolute-difference-in-bst/](https://leetcode.com/problems/minimum-absolute-difference-in-bst/)

**Example 1:**

![https://assets.leetcode.com/uploads/2021/02/05/bst1.jpg](https://assets.leetcode.com/uploads/2021/02/05/bst1.jpg)

```
Input: root = [4,2,6,1,3]
Output: 1
```

**Example 2:**

![https://assets.leetcode.com/uploads/2021/02/05/bst2.jpg](https://assets.leetcode.com/uploads/2021/02/05/bst2.jpg)

```
Input: root = [1,0,48,null,null,12,49]
Output: 1
```

**Constraints:**

- The number of nodes in the tree is in the range `[2, 100]`.
- `0 <= Node.val <= 10^5`

## 题目大意

给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

## 解题思路

- 本题和第 530 题完全相同。解题思路见第 530 题。

## 代码

```go
package leetcode

import (
	"math"

	"github.com/halfrost/leetcode-go/structures"
)

// TreeNode define
type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func minDiffInBST(root *TreeNode) int {
	res, nodes := math.MaxInt16, -1
	dfsBST(root, &res, &nodes)
	return res
}

func dfsBST(root *TreeNode, res, pre *int) {
	if root == nil {
		return
	}
	dfsBST(root.Left, res, pre)
	if *pre != -1 {
		*res = min(*res, abs(root.Val-*pre))
	}
	*pre = root.Val
	dfsBST(root.Right, res, pre)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
```