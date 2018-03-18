package main

func qsort(a []int, left, right int)  {
	if left >= right{
		return
	}
	val := a[left]
	k := left
	// 确定val 所在的位置
	for i:=left + 1; i<=right;i++{
		if a[i] < val{
			a[k] = a[i]
			a[i] = a[k+1]
			k++
		}
	}
	a[k] = val
	qsort(a, left, k-1)
	qsort(a, k+1, right)
}

func main() {
}
