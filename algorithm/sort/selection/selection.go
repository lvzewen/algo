// go run selection.go
package main

import "fmt"

func main()  {
	arr := selectionSort([]int{1, 3, 4, 2, 5})
	fmt.Printf("selection sort result %v", arr)
}

func selectionSort(arr []int) []int {
	length := len(arr)
	var minIndex int

	for i := range arr {
		minIndex = i

		for j := i + 1; j < length; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}