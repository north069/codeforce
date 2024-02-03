package main

// O(n)逆元求组合数，C(n, m) =  (n! * invM * inv(N-M)) % mod
//func CnmMod(n, m, mod int) int {
//	nP, mP, nmP := 1, 1, 1
//	for i := 2; i <= n; i++ {
//		nP = nP * i % mod
//		if i <= m {
//			mP = mP * i % mod
//		}
//		if i <= n-m {
//			nmP = nmP * i % mod
//		}
//	}
//	mP = quickPow(mP, mod-2, mod)
//	nmP = quickPow(nmP, mod-2, mod)
//	return nP * mP % mod * nmP % mod
//}
//

// O(n2)求组合数，Cnm = Cn-1,m-1 + Cn-1,m, dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
//func Cnm(n, m int) int {
//	dp := make([]int, m+1)
//	dp[0] = 1
//	for i := 0; i < n; i++ {
//		for j := m; j > 0; j-- {
//			dp[j] += dp[j-1]
//		}
//	}
//	return dp[m]
//}

//// 求数的二进制位数
//func getBit(v int64) int {
//	if v == 0 {
//		return 1
//	}
//	res := 0
//	for v > 0 {
//		res++
//		v /= 2
//	}
//	return res
//}
//
//// 判断是否是质数
//func isPrime(num int) bool {
//	if num <= 1 {
//		return false
//	}
//	for i := 2; i*i <= num; i++ {
//		if num%i == 0 {
//			return false
//		}
//	}
//	return true
//}
//
//func abs(i int) int {
//	if i < 0 {
//		return -i
//	}
//	return i
//}

//func gcd(a, b int) int {
//	if b == 0 {
//		return a
//	} else {
//		return gcd(b, a%b)
//	}
//}

//// 快速幂
//func quickPow(a, n int, mod int) int {
//	res := 1
//	for n > 0 {
//		if n&1 > 0 {
//			res *= a
//			res %= mod
//		}
//		a = a * a % mod
//		n >>= 1
//	}
//	return res % mod
//}
//
//// 拓展欧几里得
//func exgcd(a, b int) (gcd, x, y int) {
//	if b == 0 {
//		return a, 1, 0
//	}
//	gcd, y, x = exgcd(b, a%b)
//	y -= a / b * x
//	return
//}
//func exgcd(a, b int) (int, int, int) {
//	if b == 0 {
//		return a, 1, 0
//	}
//	g, x1, y1 := exgcd(b, a%b)
//	return g, y1, x1-a/b*y1
//}
//
//// 求逆元
//func invM(a, m int) int {
//	g, x, _ := exgcd(a, m)
//	if g != 1 && g != -1 {
//		return -1
//	}
//	return (x%m + m) % m
//}
//
//// 按顺序生成1到1e9+1的回文数
//func getPalArr() []int {
//	pal := make([]int, 0, 109999)
//	for base := 1; base <= 10000; base *= 10 {
//		for i := base; i < base*10; i++ {
//			x := i
//			for t := i / 10; t > 0; t /= 10 {
//				x = x*10 + t%10
//			}
//			pal = append(pal, x)
//		}
//		if base <= 1000 {
//			for i := base; i < base*10; i++ {
//				x := i
//				for t := i; t > 0; t /= 10 {
//					x = x*10 + t%10
//				}
//				pal = append(pal, x)
//			}
//		}
//	}
//	pal = append(pal, 1_000_000_001)
//	return pal
//}
