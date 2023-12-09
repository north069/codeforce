package main

//查找大于等于x中最小的一个，即x或x的后继，这种情况下取不到a[r]，可以把r设为越界无解值
func searchMax(a []int, x int) int {
	l, r := 0, len(a)-1
	for l < r {
		mid := (l + r) >> 1
		if a[mid] >= x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return a[l]
}

//查找小于等于x中最大的一个，即x或x的前驱，这种情况下取不到a[l]，可以把l设为越界无解值
//满足条件里最大的一个
func searchMin(a []int, x int) int {
	l, r := 0, len(a)
	for l < r {
		mid := (l + r + 1) >> 1
		if a[mid] <= x {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return a[l]
}

// >x 等价于 >= x+1
// >x得到下标i  i-1对应的就是<=x
//注意：使用sort.Search时，找到的是满足条件（即函数f返回值为true）的最小的索引。
