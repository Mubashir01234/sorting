package main

import (
	"fmt"
	"math/rand"
	"mubashir/sorting"
	"time"
)

func main() {
	array := make([]int, 30)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 30; i++ {
		array[i] = rand.Intn(100)
	}
	fmt.Println("Unsorted Array: ", array)
	faster := Sort(array)
	fmt.Println(faster)

}

func Sort(array []int) string {
	sortedData, bubbleSortTime := sorting.BubbleSort(array)
	fmt.Printf("\nBubble sorted: %v \nCompile Time for Bubble sort is %s\n", sortedData, bubbleSortTime)

	sortedData, insertionSortTime := sorting.InsertionSort(array)
	fmt.Printf("\nInsertion sorted: %v \nCompile Time for Insertion sort is %s\n", sortedData, insertionSortTime)

	sortedData, selectionSortTime := sorting.SelectionSort(array)
	fmt.Printf("\nSelection sorted: %v \nCompile Time for Selection sort is %s\n", sortedData, selectionSortTime)

	sortedData, quickSortTime := sorting.QuickSort(array)
	fmt.Printf("\nQuick sorted: %v \nCompile Time for Quick sort is %s\n\n", sortedData, quickSortTime)

	faster := ComparingSortTime(bubbleSortTime, insertionSortTime, selectionSortTime, quickSortTime)
	return faster
}

func ComparingSortTime(bubbleSortTime time.Duration, insertionSortTime time.Duration, selectionSortTime time.Duration, quickSortTime time.Duration) string {
	if bubbleSortTime < insertionSortTime && bubbleSortTime < selectionSortTime && bubbleSortTime < quickSortTime {
		return fmt.Sprintln("*** Bubble sort is faster than others ***")
	} else if insertionSortTime < bubbleSortTime && insertionSortTime < selectionSortTime && insertionSortTime < quickSortTime {
		return fmt.Sprintln("*** Insertion sort is faster than others ***")
	} else if selectionSortTime < bubbleSortTime && selectionSortTime < insertionSortTime && selectionSortTime < quickSortTime {
		return fmt.Sprintln("*** Selection sort is faster than others ***")
	} else if quickSortTime < bubbleSortTime && quickSortTime < selectionSortTime && quickSortTime < insertionSortTime {
		return fmt.Sprintln("*** Quick sort is faster than others ***")
	}
	return ""
}
