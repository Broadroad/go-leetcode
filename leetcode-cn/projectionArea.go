/*
在 N * N 的网格中，我们放置了一些与 x，y，z 三轴对齐的 1 * 1 * 1 立方体。

每个值 v = grid[i][j] 表示 v 个正方体叠放在单元格 (i, j) 上。

现在，我们查看这些立方体在 xy、yz 和 zx 平面上的投影。

投影就像影子，将三维形体映射到一个二维平面上。

在这里，从顶部、前面和侧面看立方体时，我们会看到“影子”。

返回所有三个投影的总面积。

Solution:
三方面：侧面，前面，顶部

*/

func projectionArea(grid [][]int) int {
	countA := 0
	countB := 0
	countC := 0

	for i := 0; i < len(grid); i++ {
		maxQian := 0
		for j := 0; j < len(grid[i]); j++ {
			// 顶部
			if grid[i][j] != 0 {
				countA++
			}

			// 前面
			if grid[i][j] > maxQian {
				maxQian = grid[i][j]
			}
		}
		countB += maxQian
	}

	// 侧面
	xlen := len(grid)
	lie := len(grid[0])
	for i := 0; i < lie; i++ {
		maxCe := 0
		for j := 0; j < xlen; j++ {
			if grid[j][i] > maxCe {
				maxCe = grid[j][i]
			}
		}
		countC += maxCe
	}

	return countA + countB + countC
}