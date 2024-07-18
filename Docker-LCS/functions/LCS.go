package functions

// max returns the larger of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// LCS calculates the length of the longest common subsequence between two strings X and Y.
func LCS_length_only(X, Y string) int {
	m := len(X)
	n := len(Y)

	// Create a 2D array dp with dimensions (m+1) x (n+1).
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Fill the dp array.
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if X[i-1] == Y[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

// LCS calculates the length of the longest common subsequence (LCS) and the LCS itself between two strings X and Y.
func LCS_with_string(X, Y string) (int, string) {
	m := len(X)
	n := len(Y)

	// Create a 2D array dp with dimensions (m+1) x (n+1).
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Create a 2D array to store the LCS.
	lcs := make([][]string, m+1)
	for i := range lcs {
		lcs[i] = make([]string, n+1)
	}

	// Fill the dp and lcs arrays.
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if X[i-1] == Y[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				lcs[i][j] = lcs[i-1][j-1] + string(X[i-1])
			} else {
				if dp[i-1][j] >= dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
					lcs[i][j] = lcs[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
					lcs[i][j] = lcs[i][j-1]
				}
			}
		}
	}

	// The length of LCS is dp[m][n].
	length := dp[m][n]
	// The LCS itself is lcs[m][n].
	lcsStr := lcs[m][n]

	return length, lcsStr
}
