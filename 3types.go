package main

import("fmt")

func add(x,y float32) float32{
	return x+y
}

func multiple(a,b string) (string, string){
	return a,b
}

func main() {


	var a int = 62
	var b float64 = float64(a)

	x := a
	fmt.Println("x is ",x)
	fmt.Println("a is ",a)
	fmt.Println("b is ",b)
}



/*
func add(x,y float32) float32{
	return x+y
}

func multiple(a,b string) (string, string){
	return a,b
}

func main() {


	w1, w2 := "Hey", "there"

	fmt.Println(multiple(w1,w2))	
}
*/

/*
func add(x float64, y float64) float64{
	return x+y
}


func main() {

	var num1 float64 = 5.6
	var num2 float64 = 9.5

	fmt.Println(add(num1,num2))
	
}

*/