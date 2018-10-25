package main

import (
	"fmt"
	"testGoProject/learn"
	"testGoProject/oop/employee"
	"testGoProject/oop/manager"
	"testGoProject/structs/computer"
)

func main() {
	//trygoroutines.Say()
	//trygoroutines.Sum()
	//TryBufferChannel()
	//tryFibonaci()
	//findInData()
	//learn.TryMaps()
	//learn.TryStruct()
	//tryMakeSpec()
	//learn.TryInterface1()
	//learn.TryInterface2()
	//learn.TryGoConcurrency()
	//learn.TryGoBufferChannel()
	//learn.TryLearnWorkerPool()
	//learn.TrySelect()
	//learn.TryMutex()
	//tryDataFromPackage()
	//learn.TryUsingComposition()
	//learn.TryStackOfDefer()

	//error handling example
	learn.TryErrorHandling()
	learn.TryDeferAndPanic()
	learn.TryFirstClassFunction()
	learn.TryReflectionData()
}

func tryDataFromPackage(){
	e := employee.Employee{
		FirstName: "Sam",
		LastName: "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()

	manager1 := manager.New("Sam", "Adolf", 30, 20)
	manager1.LeavesRemaining()
}

func tryMakeSpec(){
	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	fmt.Println("Spec : ", spec)
}

//try variadic functions
func find(num int, nums ...int){
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums{
		if v == num{
			fmt.Println(num, "found at index", i, "in" , nums)
			found = true
		}
	}

	if !found{
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func findInData(){
	find(20, 100, 300, 400, 20)
	find(200, 2000, 10)

	nums := []int{80, 90, 95}
	find(89, nums...)
	change([]string{"hello", "there"}...)
}

func change(s ...string){
	s[0] = "go"
	s = append(s, "playground")
	fmt.Println(s)
}

func TryBufferChannel(){
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("value from channel : ", <- ch)
	fmt.Println("value from channel : ", <- ch)
}

func fibonacci(n int,c chan int){
	x, y := 0, 1
	for i := 0; i < n; i++{
		c <- x
		x, y = y, x + y
	}
	close(c)
}

func tryFibonaci(){
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c{
		fmt.Println(i)
	}
}