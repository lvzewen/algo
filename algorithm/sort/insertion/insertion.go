// go run insertion.go

package main

import "fmt"

func main()  {
	arr := insertionSort([]int{1, 4, 5, 6, 3})
	fmt.Printf("result %v", arr)
}

func insertionSort(arr []int) []int {
	for i := range arr {
		perIndex := i - 1
		current := arr[i]

		for perIndex >= 0 && arr[perIndex] > current {
			arr[perIndex + 1] = arr[perIndex]
			perIndex --
		}
		arr[perIndex + 1] = current
	}
	return arr
}

