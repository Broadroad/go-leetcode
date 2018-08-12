package main

func valid(r, c, R, C int) bool {
	return 0 <= r < R && 0 <= c < C
}

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	rows := make([][]int)
	rows = append(rows, []int{r0, c0})
	step := 0

	for len(rows) < R*C {
		step += 1
		for i := 0; i < step; i++ {
			c0++
			if valid(r0, c0, R, C) {
				rows = append(rows, []int{r0, c0})
			}
		}

		for i := 0; i < step; i++ {
			r0++
			if valid(r0, c0, R, C) {
				rows = append(rows, []int{r0, c0})
			}
		}

		step += 1
		for i := 0; i < step; i++ {
			c0--
			if valid(r0, c0, R, C) {
				rows = append(rows, []int{r0, c0})
			}
		}

		for i := 0; i < step; i++ {
			r0--
			if valid(r0, c0, R, C) {
				rows = append(rows, []int{r0, c0})
			}
		}

		return rows

	}
}
