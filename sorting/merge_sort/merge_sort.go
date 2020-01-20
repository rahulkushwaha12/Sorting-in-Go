package main

import (
	"fmt"
)

/*
5
25, 17, 31, 13, 2

sorted array  [2 13 17 25 31]

Worst Case Time Complexity [ Big-O ]: O(n log(n))
Best Case Time Complexity [Big-omega]: O(n log(n))
Average Time Complexity [Big-theta]: O(n log(n))
*/
func main() {
	var n int
	fmt.Println("enter the no of inputs")
	fmt.Scanln(&n)
	arr := make([]int, n)
	fmt.Println("enter the numbers >")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	mergeSort(&arr, 0, n-1)
	fmt.Println("sorted array ", arr)
}
func mergeSort(arr *[]int, s, e int) {
	if s < e {
		mid := (s + e) / 2
		mergeSort(arr, 0, mid)
		mergeSort(arr, mid+1, e)
		merge(arr, s, mid, e)
	}
}

func merge(a *[]int, s, mid, e int) {
	arr := *a
	l1 := mid - s + 1
	l2 := e - mid
	arr1 := make([]int, l1)
	arr2 := make([]int, l2)
	copy(arr1, arr[s:mid+1])
	copy(arr2, arr[mid+1:e+1])

	var i, j, k int
	k = s
	for i < l1 && j < l2 {
		if arr1[i] <= arr2[j] {
			arr[k] = arr1[i]
			i++
		} else {
			arr[k] = arr2[j]
			j++
		}
		k++
	}
	for i < l1 {
		arr[k] = arr1[i]
		i++
		k++
	}
	for j < l2 {
		arr[k] = arr2[j]
		j++
		k++
	}
}
