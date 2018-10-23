package trygoroutines

import (
	"fmt"
	"time"
)

func _say(s string){
	for i :=0 ; i< 5; i++{
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func _sum(s []int, c chan int){
	sum := 0
	for _, v := range s{
		sum += v
	}

	c <- sum 	//send sum to c
}



//public method

func Say(){
	go _say("world")
	_say("hello")
}

func Sum(){
	s := []int{7,2,8, -9, 4, 0}
	c := make(chan int)
	go _sum(s[:len(s)/2], c)
	go _sum(s[len(s)/2:], c)
	x, y := <-c , <- c

	fmt.Println(x, y, x + y)
}

