package leetcode

import (
	"sort"

	"github.com/halfrost/leetcode-go/template"
)

// 解法一 归并排序 mergesort，时间复杂度 O(n log n)
func reversePairs(nums []int) int {
	buf := make([]int, len(nums))
	return mergesortCount(nums, buf)
}

func mergesortCount(nums, buf []int) int {
	if len(nums) <= 1 {
		return 0
	}
	mid := (len(nums) - 1) / 2
	cnt := mergesortCount(nums[:mid+1], buf)
	cnt += mergesortCount(nums[mid+1:], buf)
	for i, j := 0, mid+1; i < mid+1; i++ { // Note!!! j is increasing.
		for ; j < len(nums) && nums[i] <= 2*nums[j]; j++ {
		}
		cnt += len(nums) - j
	}
	copy(buf, nums)
	for i, j, k := 0, mid+1, 0; k < len(nums); {
		if j >= len(nums) || i < mid+1 && buf[i] > buf[j] {
			nums[k] = buf[i]
			i++
		} else {
			nums[k] = buf[j]
			j++
		}
		k++
	}
	return cnt
}

// 解法二 树状数组，时间复杂度 O(n log n)
func reversePairs1(nums []int) (cnt int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	// 离散化所有下面统计时会出现的元素
	allNums := make([]int, 0, 2*n)
	for _, v := range nums {
		allNums = append(allNums, v, 2*v)
	}
	sort.Ints(allNums)
	k := 1
	kth := map[int]int{allNums[0]: k}
	for i := 1; i < 2*n; i++ {
		if allNums[i] != allNums[i-1] {
			k++
			kth[allNums[i]] = k
		}
	}
	bit := template.BinaryIndexedTree{}
	bit.Init(k)
	for i, v := range nums {
		cnt += i - bit.Query(kth[2*v])
		bit.Add(kth[v], 1)
	}
	return
}

// 解法三 线段树，时间复杂度 O(n log n)
func reversePairs2(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	st, numsMap, indexMap, numsArray, res := template.SegmentCountTree{}, make(map[int]int, 0), make(map[int]int, 0), []int{}, 0
	numsMap[nums[0]] = nums[0]
	for _, num := range nums {
		numsMap[num] = num
		numsMap[2*num+1] = 2*num + 1
	}
	// numsArray 是 prefixSum 去重之后的版本，利用 numsMap 去重
	for _, v := range numsMap {
		numsArray = append(numsArray, v)
	}
	// 排序是为了使得线段树中的区间 left <= right，如果此处不排序，线段树中的区间有很多不合法。
	sort.Ints(numsArray)
	// 离散化，构建映射关系
	for i, n := range numsArray {
		indexMap[n] = i
	}
	numsArray = []int{}
	// 离散化，此题如果不离散化，MaxInt32 的数据会使得数字越界。
	for i := 0; i < len(indexMap); i++ {
		numsArray = append(numsArray, i)
	}
	// 初始化线段树，节点内的值都赋值为 0，即计数为 0
	st.Init(numsArray, func(i, j int) int {
		return 0
	})
	for _, num := range nums {
		res += st.Query(indexMap[num*2+1], len(indexMap)-1)
		st.UpdateCount(indexMap[num])
	}
	return res
}
