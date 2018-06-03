package main 

import (

	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string){
	defer wg.Done()
	for i:=0; i < 3; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}

func main(){
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("there")
	wg.Wait()
}

//4:00

/* 

Will print 
Doing some stuff, who knows what?
Are you done?
Done!
*/

/*
package main 

import "fmt"
	
func foo(){
	defer fmt.Println("Done!")
	defer fmt.Println("Are you done?")
	fmt.Println("Doing some stuff, who knows what?")
}

func main(){
	foo()
}
*/
