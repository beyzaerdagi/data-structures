/*
	Bubble Sort is the simplest sorting algorithm that works by repeatedly swapping the adjacent elements if they are in the wrong order.
	This algorithm is not suitable for large data sets as its average and worst-case time complexity is quite high.
	Worst and Average Case Time Complexity: O(N^2). The worst case occurs when an array is reverse sorted.
	Best Case Time Complexity: O(N). The best case occurs when an array is already sorted.
*/

package main

import "fmt"

func main() {
	var nums []int32
	nums = append(nums, 8, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(BubbleSort(nums))
}

func BubbleSort(arr []int32) []int32 {

	for i := range arr {
		noSwaps := true
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				noSwaps = false
			}
		}
		if noSwaps {
			return arr
		}
	}
	return arr
}
