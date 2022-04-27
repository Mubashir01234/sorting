package sorting

import "time"

func QuickSort(arr []int) ([]int, time.Duration) {
	timer := time.Since(time.Now())
	Sort(arr, 0, len(arr)-1)

	return arr, timer
}

func Sort(arr []int, low, high int) {
	if low < high {
		pi := Partition(arr, low, high)
		Sort(arr, low, pi-1)
		Sort(arr, pi+1, high)
	}
}

func Partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}
