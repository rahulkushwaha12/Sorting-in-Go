package main

import (
	"fmt"
)
/*
5
25, 17, 31, 13, 2

sorted array  [2 13 17 25 31]

Worst Case Time Complexity [ Big-O ]: O(n2)
Best Case Time Complexity [Big-omega]: O(n log(n))
Average Time Complexity [Big-theta]: O(n log(n))
*/
func main(){
	var n int
	fmt.Println("enter the no of inputs")
	fmt.Scanln(&n)
	arr := make([]int, n)
	fmt.Println("enter the numbers >")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	quickSort(&arr,0,n-1)
	fmt.Println("sorted array ", arr)
}

func quickSort(a *[]int, s,e int){
	if s < e {
		index := partitioning(a,s,e)
		quickSort(a,s,index-1)
		quickSort(a,index+1,e)
	}
}

func partitioning(a *[]int, s, e int )int{
	if s==e{
		return s
	}
	arr:=*a
	j := s-1
	p := arr[e]
	for i:=s;i<=e;i++{
		if arr[i]<=p{
			j++
			if arr[j] > p{
				t := arr[i]
				arr[i] = arr[j]
				arr[j] = t
			}
		}
	}
	return j
}