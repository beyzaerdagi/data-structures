/*
	Insertion sort is a simple sorting algorithm that works similar to the way you sort playing cards in your hands.
	The array is virtually split into a sorted and an unsorted part.
	Values from the unsorted part are picked and placed at the correct position in the sorted part.
	Time Complexity: O(N^2)
	Auxiliary Space: O(1)

*/

package main

import "fmt"

func main() {
	var nums []int32
	nums = append(nums, 2, 1, 9, 76, 4)
	fmt.Println(InsertionSort(nums))
}

func InsertionSort(arr []int32) []int32 {

	for i := 1; i < len(arr); i++ {
		currVal := arr[i]
		j := 0
		for j = i - 1; j >= 0 && arr[j] > currVal; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = currVal
	}
	return arr
}
