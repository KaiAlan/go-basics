package sortingalgorithms

func CountingSort(arr []int) []int {

	// find range of arr
	var n int
	for val := range arr {
		if val > n {
			n = val
		}
	}

	// find frequency of arr elements
	freq := make([]int, n+1)
	for i := 0; i < len(arr); i++ {
		freq[arr[i]]++
	}

	// find cummilative frequency
	for i := 1; i < len(arr); i++ {
		freq[i] = freq[i] + freq[i-1]
	}

	// sort the array with using cummilative freq
	sortedArr := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		sortedArr[freq[arr[i]]-1] = arr[i]
		freq[arr[i]]--
	}

	return sortedArr

}

const BaseValue = 10

func RadixSort(arr []int) []int {
	var maxElemnet int
	for _, val := range arr {
		if val > maxElemnet {
			maxElemnet = val
		}
	}

	for i := 1; maxElemnet/i > 0; i = i * BaseValue {
		arr = CountingSortForRadixSort(arr, i)
	}

	return arr
}

func CountingSortForRadixSort(arr []int, divisior int) []int {

	// find frequency of arr elements
	freq := make([]int, BaseValue)
	for i := 0; i < len(arr); i++ {
		idx := (arr[i] / divisior) % BaseValue
		freq[idx]++
	}

	// find cummilative frequency
	for i := 1; i < BaseValue; i++ {
		freq[i] = freq[i] + freq[i-1]
	}

	// sort the array with using cummilative freq
	sortedArr := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		idx := (arr[i] / divisior) % BaseValue
		sortedArr[freq[idx]-1] = arr[i]
		freq[idx]--
	}

	return sortedArr

}
