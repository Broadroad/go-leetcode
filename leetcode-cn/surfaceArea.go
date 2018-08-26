package main

import "fmt"

/*

892. 三维形体的表面积
用户通过次数 0
用户尝试次数 7
通过次数 0
提交次数 7
题目难度 Easy
在 N * N 的网格上，我们放置一些 1 * 1 * 1  的立方体。

每个值 v = grid[i][j] 表示 v 个正方体叠放在单元格 (i, j) 上。

返回结果形体的总表面积。



示例 1：

输入：[[2]]
输出：10
示例 2：

输入：[[1,2],[3,4]]
输出：34
示例 3：

输入：[[1,0],[0,2]]
输出：16
示例 4：

输入：[[1,1,1],[1,0,1],[1,1,1]]
输出：32
示例 5：

输入：[[2,2,2],[2,1,2],[2,2,2]]
输出：46

*/

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func surfaceArea(grid [][]int) int {
	sum := 0
	shouldrm := 0
	prev := 0
	// 按行来算
	for i := 0; i < len(grid); i++ {
		prev = 0
		for j := 0; j < len(grid[i]); j++ {
			sum += grid[i][j]
			minx := min(prev, grid[i][j])
			prev = grid[i][j]
			shouldrm += minx
			if grid[i][j]-1 > 0 {
				shouldrm += grid[i][j] - 1
			}
		}
	}
	fmt.Println(sum, shouldrm)

	// 按列来算
	line := len(grid)
	col := len(grid[0])
	prev = 0

	for i := 0; i < col; i++ {
		prev = 0
		for j := 0; j < line; j++ {
			minx := min(prev, grid[j][i])
			prev = grid[j][i]
			shouldrm += minx
		}
	}
	return sum*6 - 2*shouldrm
}

func main() {
	arr := [][]int{{1, 0}, {0, 2}}
	fmt.Println(surfaceArea(arr))
}
