package main

import "fmt"

// linearSearch performs a linear search through a slice of integers to
// find a specific element. It returns true if the element is found,
// false otherwise.
func linearSearch(haystack []int, needle int) bool {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}

func main() {
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	fmt.Println(linearSearch(foo, 69))
	fmt.Println(linearSearch(foo, 1124))
	fmt.Println(linearSearch(foo, 1))
	fmt.Println(linearSearch(foo, 0))
	fmt.Println(linearSearch(foo, 69420))
}
