package main

import "fmt"

func bsList(haystack []int, needle int) bool {
	low := 0
	high := len(haystack) - 1

	for low <= high {
		middle := (low + high) / 2

		if haystack[middle] == needle {
			return true
		} else if haystack[middle] < needle {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
	return false
}

func main() {
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	fmt.Println(bsList(foo, 69))
	fmt.Println(bsList(foo, 1336))
	fmt.Println(bsList(foo, 69420))
	fmt.Println(bsList(foo, 69421))
	fmt.Println(bsList(foo, 1))
	fmt.Println(bsList(foo, 0))
}
