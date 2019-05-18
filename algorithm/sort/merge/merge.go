// go run merge.go
package main

import "fmt"

func main()  {
	arr := mergeSort([]int{1, 3, 2, 4, 5, 3})
	fmt.Printf("merge sort result %v", arr)
}

func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}

	middle := length / 2		
	left := arr[0:middle]
	right := arr[middle:]

	return merge(mergeSort(left), mergeSort(right))
}


func merge (left []int, right []int) []int {
	var result []int

	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			   result = append(result, left[0])
			   left = left[1:]
		} else {
			   result = append(result, right[0])
			   right = right[1:]
		}
  	}
	
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
 	}

  	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}