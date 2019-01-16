package main

import "fmt"

func take2give2(a int, b string) (int, string) {
	fmt.Println(a)
	fmt.Println(b)
	return 20, "sheep"
}

// func main() {
// 	fmt.Println("Hello, Miguel")
// 	var foo = 0
// 	bar := 20

// 	if foo == bar {
// 		fmt.Println("foo == bar")
// 	}

// 	x, y := take2give2(foo, "sheep")
// 	fmt.Println(x)
// 	fmt.Println(y)

// 	if x == bar {
// 		fmt.Println("var == int")
// 	}
// }
