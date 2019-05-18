// go run bucket.go
package main

import "fmt"

func main()  {
	arr := bucketSort([]int{1, 3, 4, 2, 6}, 2)
	fmt.Printf("result %v", arr)
}

func bucketSort(arr []int, bucketSize int) []int{
	if len(arr) <= 1 {
		return arr
	}

	maxValue := arr[0]
	minValue := arr[0]
	for i := range arr {
		if arr[i] > maxValue {
			maxValue = arr[i]
		} else if arr[i] < minValue {
			minValue = arr[i]
		}
	}

	// 桶的初始化
	var bucketCount = (maxValue - minValue) / bucketSize + 1
	var buckets = make([][]int, bucketCount)

	for i := range arr {
		temp := (arr[i] - minValue) / bucketSize
		buckets[temp] = append(buckets[temp], arr[i])
	}

	arr = arr[0:0]
	for i := range buckets {
		insertionSort(buckets[i])
		for j := range buckets[i] {
			arr = append(arr, buckets[i][j])
		}
	}

	return arr
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