package main 

import (

	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup(){
	defer wg.Done()
	if r:= recover(); r != nil{
		fmt.Println("Recover in cleanup:", r)
	}
	
}

func say(s string){
	defer cleanup()
	for i:= 0; i < 3; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
		if i == 2 {
			panic("Oh dear, a 2")
		}
	}
}

func main(){
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("there")
	wg.Wait()
}