/*
	The Merge Sort algorithm is a sorting algorithm that is considered an example of the divide and conquer strategy.
	So, in this algorithm, the array is initially divided into two equal halves and then they are combined in a sorted manner.
	We can think of it as a recursive algorithm that continuously splits the array in half until it cannot be further divided.
	This means that if the array becomes empty or has only one element left, the dividing will stop, i.e. it is the base case to stop the recursion.
	If the array has multiple elements, we split the array into halves and recursively invoke the merge sort on each of the halves.
	Finally, when both the halves are sorted, the merge operation is applied.
	Merge operation is the process of taking two smaller sorted arrays and combining them to eventually make a larger one.
	Time Complexity: O(n log(n))

*/

package main

import "fmt"

func main() {
	var arr []int32
	arr = append(arr, 10, 24, 76, 73)

	fmt.Println(MergeSort(arr))
}

func merge(arr1, arr2 []int32) []int32 {

	var sortedArr []int32

	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] >= arr2[j] {
			sortedArr = append(sortedArr, arr2[j])
			j++
		} else {
			sortedArr = append(sortedArr, arr1[i])
			i++
		}
	}

	for i < len(arr1) {
		sortedArr = append(sortedArr, arr1[i])
		i++
	}

	for j < len(arr2) {
		sortedArr = append(sortedArr, arr2[j])
		j++
	}
	return sortedArr
}

func MergeSort(arr []int32) []int32 {

	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[0:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}
