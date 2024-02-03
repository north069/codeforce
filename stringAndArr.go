package main

// 数组原地排序去重
//func unique(nums []int) []int {
//	sort.Ints(nums)
//	res := make([]int, len(nums))
//	k := -1
//	for _, v := range nums {
//		if (k < 0 || res[k] != v) && v >= 0 {
//			k++
//			res[k] = v
//		}
//	}
//	return res[:k+1]
//}

// 返回整个数组到v的距离
// 求前缀和
//func getArrDis(nums []int, v int) int {
//	n := len(nums)
//	pre := make([]int, n+1)
//	for i := 1; i <= n; i++ {
//		pre[i] = pre[i-1] + nums[i-1]
//	}
//	getDis := func(v int) int {
//		i := sort.SearchInts(nums, v)
//		res := i*v - pre[i]
//		res += pre[n] - pre[i] - v*(n-i)
//		return res
//	}
//	return getDis(v)
//}

// 让数组相等的最小代价
//func LetNumsEqual(nums []int) int {
//	sort.Ints(nums)
//	res := 0
//	n := len(nums)
//	for i := 0; i < n; i++ {
//		res += nums[i] - nums[i/2]
//	}
//	return res
//}

// 字符串反转
//func reverse(num string) string {
//	temp := []byte(num)
//	left, right := 0, len(num)-1
//	for left < right {
//		temp[left], temp[right] = temp[right], temp[left]
//		left++
//		right--
//	}
//	return string(temp)
//}

//kmp匹配
//func buildKMPTable(pattern string) []int {
//	table := make([]int, len(pattern))
//	j := 0
//
//	for i := 1; i < len(pattern); i++ {
//		if pattern[i] == pattern[j] {
//			j++
//			table[i] = j
//		} else {
//			if j != 0 {
//				j = table[j-1]
//				i--
//			} else {
//				table[i] = 0
//			}
//		}
//	}
//
//	return table
//}
//func kmpSearch(text, pattern string) []int {
//	table := buildKMPTable(pattern)
//	m, n := len(pattern), len(text)
//	result := []int{}
//
//	i, j := 0, 0
//	for i < n {
//		if pattern[j] == text[i] {
//			i++
//			j++
//		}
//
//		if j == m {
//			result = append(result, i-j)
//			j = table[j-1]
//		} else if i < n && pattern[j] != text[i] {
//			if j != 0 {
//				j = table[j-1]
//			} else {
//				i++
//			}
//		}
//	}
//
//	return result
//}
