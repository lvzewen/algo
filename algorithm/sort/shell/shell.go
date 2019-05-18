// go run shell.go
package main

import "fmt"

func main()  {
	arr := shellSort([]int{1, 3, 4, 2, 6})
	fmt.Printf("shell sort result %v", arr)
}

func shellSort(arr []int) []int {
	length := len(arr);
	var temp int
	
	gap := 1
	for gap < length / 3 {
		gap = gap * 3 + 1
	}

	for gap > 0{
		for i := gap; i < length; i++ {
			temp = arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j + gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 3
	}
	return arr
}