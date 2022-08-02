/*
	QuickSort is a Divide and Conquer algorithm. It picks an element as a pivot and partitions the given array around the picked pivot.
	Space Complexity: O(log n)
	Average: O(n*log n)
	Worst: O(n^2)
*/

package main

import "fmt"

func main() {
	var arr []int32
	arr = append(arr, 10, 24, 76, 73)

	fmt.Println(QuickSort(arr, 0, int32(len(arr))-1))
}

func pivot(arr []int32, start, end int32) int32 {

	pivot := arr[start]
	swapIdx := start

	for i := start + 1; i < int32(len(arr)); i++ {
		if pivot > arr[i] {
			swapIdx++
			arr[i], arr[swapIdx] = arr[swapIdx], arr[i]
		}
	}
	arr[start], arr[swapIdx] = arr[swapIdx], arr[start]
	return swapIdx
}

func QuickSort(arr []int32, start, end int32) []int32 {
	if start > end {
		return arr
	}

	p := pivot(arr, start, end)
	QuickSort(arr, start, p-1)
	QuickSort(arr, p+1, end)
	return arr
}
