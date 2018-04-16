package main 

import("fmt"
		"math"
		"math/rand")

func foo(){
		fmt.Println("The square root of 4 is", math.Sqrt(4))

}

func random(){
	fmt.Println(" A number from 1-1000",rand.Intn(100))
}



func main() {
	foo()
	random()
}

/* Notes

fun roo(){
	//Do something
	// something()
	//Do something to somehting
	// something.Int(C)
}

func main(){
	roo()
	//does two things because roo did two things
}
/