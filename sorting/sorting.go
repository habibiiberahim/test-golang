package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 4, 5, 6, 8, 2}
	sortingWithVisual(arr)
	sortingWithVisualReverse(arr)
}

func findMax(a []int) (max int) {
	max = a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}

func arrayWithVisual(a []int) {
	max := findMax(a)
	for i := max; i >= 0; i-- {
		for j := 0; j < len(a); j++ {
			if i > a[j] {
				fmt.Print(" ")
			} else if i == 0 {
				fmt.Print(a[j])
			} else {
				fmt.Print("|")
			}
		}
		fmt.Println()
	}
}

func sortingWithVisual(a []int) {
	arr := a
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
			arrayWithVisual(arr)
		}
	}
}

func sortingWithVisualReverse(a []int) {
	arr := a
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
			arrayWithVisual(arr)
		}
	}

}
