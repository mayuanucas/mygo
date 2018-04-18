package problem441

import "math"

// 题目的意思其实就是从1~x层完整楼梯硬币数量加起来，要小于等于n，求最大的x。
// 说到加起来的数量，很容易想到求累加和，我们知道求累加和的公式为：
// sum = (1+x)*x/2
// 这里就是要求 sum <= n 了。我们反过来求层数x。如果直接开方来求会存在错误，必须因式分解求得准确的x值：
//
// (1+x)*x/2 <= n
// x*x + x <= 2*n
// 4*x*x + 4*x <= 8*n
// (2*x + 1)*(2*x + 1) - 1 <= 8*n
// x <= (sqrt(8*n + 1) - 1) / 2

func arrangeCoins(n int) int {
	return int((math.Sqrt(float64(8*n+1)) - 1) / 2)
}
