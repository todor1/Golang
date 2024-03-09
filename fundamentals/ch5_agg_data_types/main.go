package main

import "fmt"

// aggreagate data types: array, slice, map, struct
func main() {
	var arr [3]int
	fmt.Println(arr) // [0 0 0]
	arr = [3]int{1, 2, 3}
	fmt.Println(arr)
}
