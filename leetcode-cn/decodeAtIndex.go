func decodeAtIndex(S string, K int) string {
	n := len(S)

	dp := make([]int, n+1)
	for i, s := range S {

		if s >= '2' && s <= '9' {
			dp[i+1] = dp[i] * int(s-'0')
		} else {
			dp[i+1] = dp[i] + 1
		}
	}
	K--
	for i := n - 1; i >= 0; i-- {
		K %= dp[i+1]
		if K+1 == dp[i+1] && !(S[i] >= '2' && S[i] <= '9') {
			return string(S[i])
		}
	}

	return ""
}
