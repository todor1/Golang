package main

import "fmt"

// var x int
var x = 32

func main() {
	// println(x)
	// fmt.Println(x)
	// fmt.Printf("x is of type %T\n", x)
	y := 42.4
	// fmt.Println(y)
	// fmt.Printf("y is of type %T\n", y)
	z := "doh"
	// fmt.Println(z)
	// fmt.Printf("z is of type %T\n", z)

	// outputThat(42, 42.4, "doh")
	dpne, trueOrNot := outputThat(x, y, z)
	fmt.Println(dpne, trueOrNot)
}

func outputThat(a int, b float64, c string) (string, bool) {
	fmt.Printf("Integer = %d, float = %f, string = %s\n", a, b, c)
	return "ok", true
}
