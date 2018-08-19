func possibleBipartition(N int, dislikes [][]int) bool {
	flags := make([][]int, N)
	for i := 0; i< N;i++ {
		col := make([]int, N)
		flags[i] = col
	}
	for k, v := range dislikes{
		r := v[0]
		c := v[1]

		flags[r][c] = true
	}

	count := 0
	has := make(int[], N)
	hasnt := make(int[], N)
	for i := 1; i <= N; i++ {
		hasnt [i-1] = i
	}

	dfs(0,1,disklikes, has, hasnt, flags, count) {

	}
}

func dfs(x, y int, dislikes [][]int, has []int, hasnt []int, flags [][]int, count int) bool {
	if flags[x][y] == true {
		return false
	}

	if count == N - 1 {
		return true
	}
	
	for 
}