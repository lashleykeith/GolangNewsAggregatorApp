package main

import "fmt"
		

type car struct {
	gas_pedal uint16 // min 0 max 65535
	brake_pedal uint16
	steering_wheel int16 // -32k - +32k
	top_speed_kmh float64
}

type human struct {
	skin string
	hair string
	height string

}

func main(){
	a_car := car{gas_pedal: 22341, 
				brake_pedal: 0, 
				steering_wheel: 12561, 
				top_speed_kmh: 225.0 }

	a_human := human{
					skin: "Great",
					hair: "Good",
					height: "Tall"}
					
	fmt.Println(a_human.skin)
	fmt.Println(a_car.gas_pedal)

}