package main

import "fmt"

func main() {

	x := 20
	a := &x
	fmt.Println(a)
	fmt.Println(*a)
	*a = 4
	fmt.Println(x)
	*a = *a**a
	fmt.Println(x)
	fmt.Println(*a)

	
}

