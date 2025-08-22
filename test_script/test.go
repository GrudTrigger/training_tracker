package main

import (
	"fmt"
)

func BubbleSort(arr []int, descending bool) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		if descending {
			for j := 0; j < n-i-1; j++ {
				if arr[j] < arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					swapped = true
				}
			}
			if !swapped {
				break
			}
		} else {
			for j := 0; j < n-i-1; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					swapped = true
				}
			}
			if !swapped {
				break
			}
		}

	}
	return arr
}

// Quick Sort
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[(len(arr)-1)/2]

	var left []int
	var right []int

	for i, v := range arr {
		if i == (len(arr)-1)/2 {
			continue
		}
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	result := append(QuickSort(left), pivot)
	return append(result, QuickSort(right)...)
}

// Insertion sort
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i] // текущий элемент
		j := i - 1

		// Двигаем элементы вправо, пока они больше key
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		// Вставляем key на правильное место
		arr[j+1] = key
	}
}

func InsertSortStrings(arr []string) {
	for i := 1; i < len(arr); i++ {
		char := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > char {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = char
	}
}

func main() {
	arr := []int{33, 10, 55, 71, 29, 4, 12}
	//res := BubbleSort(arr, true)
	//fmt.Println(res)
	//arr := []string{"g", "q", "x", "z", "a", "l", "y", "k"}
	//InsertSortStrings(arr)
	//fmt.Println(arr)

	r := QuickSort(arr)
	fmt.Println(r)
}
