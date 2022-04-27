package sorting

import "time"

func SelectionSort(array []int) ([]int, time.Duration) {
	timer := time.Since(time.Now())
	var min_index int
	var temp int
	for i := 0; i < len(array)-1; i++ {
		min_index = i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min_index] {
				min_index = j
			}
		}
		temp = array[i]
		array[i] = array[min_index]
		array[min_index] = temp
	}
	return array, timer
}
