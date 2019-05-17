// 可以转换为等差数列求和公式 n = m*f + m*(m-1)/2;  
// n 为要拆分的整数，  
// m 为拆分后连续自然数个数  
// f 为连续自然数中的第一位数 

// 空间复杂度O(1) 只考虑初始化变量所占用的空间
// 时间复杂度O(nlogn) 包含最外层的遍历，与内部的求平方根

package main

import "fmt"

var (
	n,f int
	w,k,m int
)

func main()  { 
	fmt.Println("Please enter number: ")
	fmt.Scanln(&n)

	for i := 1; i < n / 2; i++ {
		f = i                  
		w = (2*f-1) * (2*f-1) + 8*n  
		k = mySqrt2(w) 
		m = k - 2*f + 1
	   
		if k*k == w && m % 2 == 0 {
			fmt.Printf("%d = ", n);
			for j := 1; j< m / 2; j++ {
				if j == 1 {
					fmt.Printf("%d ", i + j - 1);
				} else {
					fmt.Printf("+ %d ", i + j - 1);
				}
				
			}
			fmt.Printf("\n");
		 }
	}
}

// 在有序数组中，找到最后一个小于等于给定值的数
func mySqrt2(x int) int {
    low, high := 0, x
    for low <= high {
        //防止大数相加溢出
        //位运算更高效
        mid := low + (high-low)>>1
        product := mid * mid
        if product > x {
            high = mid - 1
        } else {
            if (mid == x) || (mid+1)*(mid+1) > x {
                //遍历最后一个数 || 下一个数大于目标值
                return mid
            }
            //下一个数小于等于目标值，所以mid不是最后一个数
            low = mid + 1
        }
    }
    return -1
}