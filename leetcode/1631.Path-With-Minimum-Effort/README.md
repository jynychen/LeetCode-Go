# [1631. Path With Minimum Effort](https://leetcode.com/problems/path-with-minimum-effort/)

## 题目

You are a hiker preparing for an upcoming hike. You are given `heights`, a 2D array of size `rows x columns`, where `heights[row][col]` represents the height of cell `(row, col)`. You are situated in the top-left cell, `(0, 0)`, and you hope to travel to the bottom-right cell, `(rows-1, columns-1)` (i.e., **0-indexed**). You can move **up**, **down**, **left**, or **right**, and you wish to find a route that requires the minimum **effort**.

A route's **effort** is the **maximum absolute difference** in heights between two consecutive cells of the route.

Return *the minimum **effort** required to travel from the top-left cell to the bottom-right cell.*

**Example 1:**

![https://assets.leetcode.com/uploads/2020/10/04/ex1.png](https://assets.leetcode.com/uploads/2020/10/04/ex1.png)

```
Input: heights = [[1,2,2],[3,8,2],[5,3,5]]
Output: 2
Explanation: The route of [1,3,5,3,5] has a maximum absolute difference of 2 in consecutive cells.
This is better than the route of [1,2,2,2,5], where the maximum absolute difference is 3.
```

**Example 2:**

![https://assets.leetcode.com/uploads/2020/10/04/ex2.png](https://assets.leetcode.com/uploads/2020/10/04/ex2.png)

```
Input: heights = [[1,2,3],[3,8,4],[5,3,5]]
Output: 1
Explanation: The route of [1,2,3,4,5] has a maximum absolute difference of 1 in consecutive cells, which is better than route [1,3,5,3,5].
```

**Example 3:**

![https://assets.leetcode.com/uploads/2020/10/04/ex3.png](https://assets.leetcode.com/uploads/2020/10/04/ex3.png)

```
Input: heights = [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
Output: 0
Explanation: This route does not require any effort.
```

**Constraints:**

- `rows == heights.length`
- `columns == heights[i].length`
- `1 <= rows, columns <= 100`
- `1 <= heights[i][j] <= 10^6`

## 题目大意

你准备参加一场远足活动。给你一个二维 `rows x columns` 的地图 `heights` ，其中 `heights[row][col]` 表示格子 `(row, col)` 的高度。一开始你在最左上角的格子 `(0, 0)` ，且你希望去最右下角的格子 `(rows-1, columns-1)` （注意下标从 0 开始编号）。你每次可以往 上，下，左，右 四个方向之一移动，你想要找到耗费 体力 最小的一条路径。一条路径耗费的 体力值 是路径上相邻格子之间 高度差绝对值 的 最大值 决定的。请你返回从左上角走到右下角的最小 体力消耗值 。

## 解题思路

- 此题和第 778 题解题思路完全一致。在第 778 题中求的是最短连通时间。此题求的是连通路径下的最小体力值。都是求的最小值，只是 2 个值的意义不同罢了。
- 按照第 778 题的思路，本题也有多种解法。第一种解法是 DFS + 二分。先将题目变换一个等价问法。题目要求找到最小体力消耗值，也相当于问是否存在一个体力消耗值 x，只要大于等于 x，一定能连通。利用二分搜索来找到这个临界值。体力消耗值是有序的，此处满足二分搜索的条件。题目给定柱子高度是 [1,10^6]，所以体力值一定在 [0,10^6-1] 这个区间内。判断是否取中值的条件是用 DFS 或者 BFS 搜索 (0,0) 点和 (N-1, N-1) 点之间是否连通。时间复杂度：O(mnlogC)，其中 m 和 n 分别是地图的行数和列数，C 是格子的最大高度。C 最大为 10^6，所以 logC 常数也很小。空间复杂度 O(mn)。
- 第二种解法是并查集。将图中所有边按照权值从小到大进行排序，并依次加入并查集中。直到加入一条权值为 x 的边以后，左上角到右下角连通了。最小体力消耗值也就找到了。注意加入边的时候，只加入 `i-1` 和 `i` ，`j-1` 和 `j` 这 2 类相邻的边。因为最小体力消耗意味着不走回头路。上下左右四个方向到达一个节点，只可能从上边和左边走过来。从下边和右边走过来肯定是浪费体力了。时间复杂度：O(mnlog(mn))，其中 m 和 n 分别是地图的行数和列数，图中的边数为 O(mn)。空间复杂度 O(mn)，即为存储所有边以及并查集需要的空间。

## 代码

```go
package leetcode

import (
	"sort"

	"github.com/halfrost/leetcode-go/template"
)

var dir = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

// 解法一 DFS + 二分
func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}
	low, high := 0, 1000000
	for low < high {
		threshold := low + (high-low)>>1
		if !hasPath(heights, visited, 0, 0, threshold) {
			low = threshold + 1
		} else {
			high = threshold
		}
		for i := range visited {
			for j := range visited[i] {
				visited[i][j] = false
			}
		}
	}
	return low
}

func hasPath(heights [][]int, visited [][]bool, i, j, threshold int) bool {
	n, m := len(heights), len(heights[0])
	if i == n-1 && j == m-1 {
		return true
	}
	visited[i][j] = true
	res := false
	for _, d := range dir {
		ni, nj := i+d[0], j+d[1]
		if ni < 0 || ni >= n || nj < 0 || nj >= m || visited[ni][nj] || res {
			continue
		}
		diff := abs(heights[i][j] - heights[ni][nj])
		if diff <= threshold && hasPath(heights, visited, ni, nj, threshold) {
			res = true
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 解法二 并查集
func minimumEffortPath1(heights [][]int) int {
	n, m, edges, uf := len(heights), len(heights[0]), []edge{}, template.UnionFind{}
	uf.Init(n * m)
	for i, row := range heights {
		for j, h := range row {
			id := i*m + j
			if i > 0 {
				edges = append(edges, edge{id - m, id, abs(h - heights[i-1][j])})
			}
			if j > 0 {
				edges = append(edges, edge{id - 1, id, abs(h - heights[i][j-1])})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].diff < edges[j].diff })
	for _, e := range edges {
		uf.Union(e.v, e.w)
		if uf.Find(0) == uf.Find(n*m-1) {
			return e.diff
		}
	}
	return 0
}

type edge struct {
	v, w, diff int
}
```