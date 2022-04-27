package sorting

import "time"

func InsertionSort(arr []int) ([]int, time.Duration) {
	timer := time.Since(time.Now())
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr, timer
}
