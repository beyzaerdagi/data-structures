/*
	The selection sort algorithm sorts an array by repeatedly finding the minimum element from unsorted part and putting it at the beginning.
	The algorithm maintains two subarrays in a given array.
	Time Complexity: O(N^2)
*/

package main

import "fmt"

func main() {
	var nums []int32
	nums = append(nums, 81, 1, 2, 3, 49, 54, 6, 7, 18, 33, 22, 43)
	fmt.Println(SelectionSort(nums))
}

func SelectionSort(arr []int32) []int32 {

	for i := range arr {
		smallest := i
		for j := i + 1; j < len(arr); j++ {
			if arr[smallest] > arr[j] {
				smallest = j
			}
		}
		if smallest != i {
			arr[i], arr[smallest] = arr[smallest], arr[i]
		}
	}
	return arr
}
