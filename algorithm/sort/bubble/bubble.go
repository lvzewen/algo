// go run bubble.go
package main

import "fmt"

func main()  {
	arr := bubbleSort([]int{1, 3, 4, 2, 6})
	fmt.Printf("result %v", arr)
}

func bubbleSort(arr []int) []int {
	length := len(arr)

	for i := 0; i < length; i++ {
		for j := 0; j < length - i - 1; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}