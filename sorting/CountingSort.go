package main

// func CountingSort(arr []int) []int {

// 	// find range of arr
// 	var n int
// 	for val := range arr {
// 		if val > n {
// 			n = val
// 		}
// 	}

// 	// find frequency of arr elements
// 	freq := make([]int, n+1)
// 	for i := 0; i < len(arr); i++ {
// 		freq[arr[i]]++
// 	}

// 	// find cummilative frequency
// 	for i := 1; i < len(arr); i++ {
// 		freq[i] = freq[i] + freq[i-1]
// 	}

// 	// sort the array with using cummilative freq
// 	sortedArr := make([]int, len(arr))
// 	for i := len(arr) - 1; i >= 0; i-- {
// 		sortedArr[freq[arr[i]]-1] = arr[i]
// 		freq[arr[i]]--
// 	}

// 	return sortedArr

// }
