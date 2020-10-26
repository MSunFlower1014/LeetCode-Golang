package main

import "fmt"

func main() {
	test := []int{1, 2, 4, 3, 9, 1}
	quickSort(test, 0, len(test)-1)
	fmt.Printf("%v", test)
}

/*
快速排序
*/
func quickSort(s []int, left int, right int) {
	/*
		任何时刻指向的都是小于等于s[left]的位置,遇到一个小于left的explodeIndex就会++并且交换，当a[i】<= a[left]的时候，要末和自己交换，要么和第一个大于s【left】的交换
	*/
	if left >= right {
		return
	}
	changeIndex := left //0
	for i := left + 1; i <= right; i++ {
		if s[i] <= s[left] {
			changeIndex++
			s[i], s[changeIndex] = s[changeIndex], s[i]
		}
	}
	s[changeIndex], s[left] = s[left], s[changeIndex]
	quickSort(s, left, changeIndex-1)
	quickSort(s, changeIndex+1, right)
}

/*
归并排序
*/
func mergeSort(left int, right int, arr []int) []int {

	if len(arr) == 1 {
		return arr
	}
	mid := (left + right) / 2

	a1 := mergeSort(left, mid, arr[0:mid+1-left])
	a2 := mergeSort(mid+1, right, arr[mid+1-left:])

	len1 := len(a1)
	len2 := len(a2)

	temp := make([]int, 0)
	for i, j := 0, 0; i < len1 || j < len2; {
		if i == len1 {
			temp = append(temp, a2[j])
			j++
			continue
		}
		if j == len2 {
			temp = append(temp, a1[i])
			i++
			continue
		}
		if a1[i] < a2[j] {
			temp = append(temp, a1[i])
			i++
		} else {
			temp = append(temp, a2[j])
			j++
		}
	}
	return temp
}
