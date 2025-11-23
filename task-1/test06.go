package main

import (
	"fmt"
	"sort"
)

/*
56. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区间的起始位置进行排序，
然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
*/

func merge(intervals [][]int) [][]int {
	// 如果区间数量小于等于1，直接返回
	if len(intervals) <= 1 {
		return intervals
	}

	// 1. 按照区间的起始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 2. 初始化结果切片，先加入第一个区间
	result := [][]int{intervals[0]}

	// 3. 遍历排序后的区间数组
	for i := 1; i < len(intervals); i++ {
		// 获取当前区间和结果中的最后一个区间
		current := intervals[i]
		last := result[len(result)-1]

		// 4. 检查是否有重叠
		if current[0] <= last[1] {
			// 有重叠，合并区间（取结束位置的最大值）
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			// 没有重叠，将当前区间添加到结果中
			result = append(result, current)
		}
	}

	return result
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	intervals = merge(intervals)
	fmt.Println(intervals)

}
