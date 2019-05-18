// go run quick.go
package main

import "fmt"

func main()  {
	arr := quickSort([]int{1, 3, 4, 2, 5})
	fmt.Printf("quick sort result %v", arr)
}

func quickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
	if left < right {
		   partitionIndex := partition(arr, left, right)
		   _quickSort(arr, left, partitionIndex-1)
		   _quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index ++
		}
	}
	swap(arr, pivot, index-1)
	return index - 1
}


func swap (arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}