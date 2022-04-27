package sorting

import "time"

func BubbleSort(array []int) ([]int, time.Duration) {
	timer := time.Since(time.Now())
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array, timer
}
