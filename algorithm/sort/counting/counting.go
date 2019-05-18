// go run counting

package main

import "fmt"

func main()  {
	arr := countingSort([]int{1, 3, 4, 2, 6}, 6)
	fmt.Printf("result %v", arr)
}

func countingSort(arr []int, maxValue int) []int{
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen) // 初始为0的数组
	
	sortedIndex := 0
	length := len(arr)
	   
	for i := 0; i < length; i++ {
		bucket[arr[i]] ++
	}

	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			arr[sortedIndex] = j
			sortedIndex ++
			bucket[j] --
		}
     }

	return arr
}