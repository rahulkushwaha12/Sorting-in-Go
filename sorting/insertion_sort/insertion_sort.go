package main

import (
	"fmt"
)

/*
5
25, 17, 31, 13, 2

sorted array  [2 13 17 25 31]

Worst Case Time Complexity [ Big-O ]: O(n2)
Best Case Time Complexity [Big-omega]: O(n)
Average Time Complexity [Big-theta]: O(n2)
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
	for i := 1; i < n; i++ {
		valToInsert := arr[i]
		k := i
		for k > 0 && arr[k-1] > valToInsert {
			arr[k] = arr[k-1]
			k--
		}
		if k != i {
			arr[k] = valToInsert
		}
	}
	fmt.Println("sorted array ", arr)
}
