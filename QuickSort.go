package functions

// QuickSort sorts an array using the quicksort algorithm.
func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	// Call the recursive quicksort function with initial low and high indices.
	quickSort(arr, 0, len(arr)-1)
}

// quickSort is the recursive function for the quicksort algorithm.
func quickSort(arr []int, low, high int) {
	if low < high {
		// Partition the array and get the pivot index.
		p := partition(arr, low, high)
		// Recursively sort elements before and after partition.
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

// partition rearranges the elements in the array and returns the pivot index.
func partition(arr []int, low, high int) int {
	// Choose the pivot element.
	pivot := arr[high]
	// Index of the smaller element.
	i := low - 1

	for j := low; j < high; j++ {
		// If the current element is smaller than or equal to the pivot.
		if arr[j] <= pivot {
			i++
			// Swap arr[i] and arr[j].
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// Swap arr[i+1] and arr[high] (or pivot).
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
