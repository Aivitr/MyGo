package main

import "math"

func fibIterative(n int) int {
	if n < 0 {
		panic("n 不能为负数")
	}
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b // 滚动更新：a→F(i+1), b→F(i+2)
	}
	return a
}

func fibRecursive(n int) int {
	if n < 0 {
		panic("n 不能为负数")
	}
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibRecursive(n-1) + fibRecursive(n-2)
	}
}

func fibMemo(n int) int {
	var fibCache = make(map[int]int)
	if n < 0 {
		panic("n 不能为负数")
	}
	if val, ok := fibCache[n]; ok {
		return val
	}
	switch n {
	case 0:
		fibCache[n] = 0
	case 1:
		fibCache[n] = 1
	default:
		fibCache[n] = fibMemo(n-1) + fibMemo(n-2)
	}
	return fibCache[n]
}

func fibDP(n int) int {
	if n < 0 {
		panic("n 不能为负数")
	}
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

//F(n) ≈ (φⁿ - ψⁿ)/√5
func fibFormula(n int) int {
	if n < 0 {
		panic("n 不能为负数")
	}
	phi := (1 + math.Sqrt(5)) / 2
	psi := (1 - math.Sqrt(5)) / 2
	return int(math.Round((math.Pow(phi, float64(n)) - math.Pow(psi, float64(n))) / math.Sqrt(5)))
}

//F(2k) = F(k) * [ 2*F(k+1) - F(k) ]
//F(2k+1) = F(k+1)² + F(k)²
func fibFormulaFast(n int) int64 {
	if n < 0 {
		panic("n 不能为负数")
	}
	var fast func(int) (int64, int64)
	fast = func(k int) (int64, int64) {
		if k == 0 {
			return 0, 1
		}
		a, b := fast(k / 2)
		c := a * (2*b - a)
		d := a*a + b*b
		if k%2 == 0 {
			return c, d
		}
		return d, c + d
	}
	res, _ := fast(n)
	return res
}
