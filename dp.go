package main

import "strconv"

//数位dp模板
func countSpecialNumbers(n int) int {
	str := strconv.Itoa(n)
	m := len(str)
	memo := make([][1 << 10]int, m)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var f func(i, mask int, isLimit bool, isNum bool) int
	f = func(i, mask int, isLimit, isNum bool) (ans int) {
		if i == m {
			if isNum {
				return 1
			}
			return
		}
		//保存非边界的状态
		if !isLimit && isNum {
			p := &memo[i][mask]
			if *p >= 0 {
				return *p
			}
			defer func() {
				*p = ans
			}()
		}
		if !isNum {
			ans += f(i+1, mask, false, false)
		}
		d := 0
		if !isNum {
			d = 1
		}
		up := 9
		if isLimit {
			up = int(str[i] - '0')
		}
		for ; d <= up; d++ {
			if mask>>d&1 == 0 {
				ans += f(i+1, mask|1<<d, isLimit && d == up, true)
			}
		}
		return
	}

	return f(0, 0, true, false)
}

//枚举子集的写法 O(3^n)
//for i := chosen; i > 0; i = (i - 1) & chosen {
//	if bits.OnesCount(uint(i)) == k {
//		ans = min(ans, 1+dfs(mask-i))
//	}
//}
