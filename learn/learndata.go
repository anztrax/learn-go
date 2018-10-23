package learn

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Address struct {
	city, state string
}

type Employee struct{
	firstName string
	lastName string
	currency string
	age, salary int
	address Address
}

func(e Employee) displaySalary(){
	fmt.Printf("Salary of %s %s # age : %d is %s%d\n", e.firstName, e.lastName, e.age , e.currency, e.salary)
}
func(e Employee) displayAddress(){
	fmt.Printf("full address : %s %s\n", e.address.city, e.address.state)
}
func(e Employee) setFirstName(firstName string){
	e.firstName = firstName
}
func(e *Employee) setAge(newAge int){
	e.age = newAge
}

type image struct{
	data map[int]int
}


//global function
func TryMaps(){
	var personSallary map[string]int
	if personSallary == nil{
		fmt.Println("map is nil, going to make one .")
		personSallary = make(map[string]int)
		personSallary["steve"] = 12000
		personSallary["jamie"] = 15000
		personSallary["mike"] = 9000
		fmt.Println("person salary map contents:", personSallary)

		personSalary2 := map[string]int{
			"steve" : 12000,
			"jamie" : 15000,
			"testing" : 30000,
		}
		delete(personSalary2, "testing")
		personSalary2["mike"] = 9000
		value, ok := personSalary2["jamie"]
		if ok{
			fmt.Println("jamie salary", value)
		}else{
			fmt.Println("can't find jamie salary")
		}
		allPersonSalary(personSalary2)
	}
}

func TryStruct(){
	emp1 := Employee{
		age: 10,
		firstName:"david",
		lastName: "gueta",
		address: Address {
			city: "jakarta",
			state: "DKI. JAKARTA",
		},
	}
	emp1.setAge(30)
	emp1.setFirstName("nutella")
	emp1.displaySalary()
	emp1.displayAddress()

	//anonymous struct
	emp3 := struct{
		firstName, lastName string
		age, salary int
	}{
		firstName: "hello",
		lastName: "there",
	}
	fmt.Println("emp3 : ", emp3)

	//struct are comparable if fields on struct is comparable
	address1 := Address{
		state: "jakarta",
		city: "jakarta",
	}
	address2 := Address{
		state: "jakarta",
		city: "jakarta",
	}
	fmt.Println("address are equals : ", address1 == address2)


	image1 := image{
		data: map[int]int{
			0 : 155,
		},
	}
	image2 := image{
		data : map[int]int{
			0 : 200,
		},
	}
	fmt.Printf("image 1 : %V, %V\n", image1, image2)
}

type VowelFinder interface {
	FindVowels() []rune
}

type MyString string

func(ms MyString) FindVowels() [] rune{
	var vowels []rune
	for _, rune := range ms{
		if rune == 'a' || rune == 'i' || rune == 'u' || rune == 'e' || rune == 'o'{
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

type SalaryCalculator interface {
	CalculateSalary() int
}
type Permanent struct{
	empId int
	basicPay int
	pf int
}
type Contract struct{
	empId int
	basicPay int
}

func(p Permanent) CalculateSalary() int{
	return p.basicPay + p.pf
}
func(c Contract) CalculateSalary() int{
	return c.basicPay
}

func totalExpensesForSalary(s []SalaryCalculator){
	expense := 0
	for _, v := range s{
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expenses Per Month $%d\n", expense)
}

//interface{} <- general object
func describe(i interface{}){
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func TryInterface1(){
	name := MyString("Sam Anderson")
	var v VowelFinder = name
	fmt.Printf("string %s , Vowels are %c\n\n", name, v.FindVowels())

	pe1 := Permanent{basicPay: 100, empId: 01, pf: 200}
	pe2 := Permanent{basicPay: 100, empId: 02, pf: 300}
	co1 := Contract{empId: 03, basicPay: 300}
	totalExpensesForSalary([]SalaryCalculator{pe1, pe2, co1})

	s := "Hello world"
	describe(s)

	struct1 := struct {
		name string
	}{
		name : "testing gan",
	}
	describe(struct1)

}

type SalaryCalculator2 interface {
	DisplaySalary2()
}
type LeaveCalculator interface {
	CalculatorLeavesLeft() int
}

type EmployeeOperations interface {
	SalaryCalculator2
	LeaveCalculator
}

type Employee2 struct{
	firstName string
	lastName string
	basicPay int
	totalLeaves int
	leavesTaken int
}
func(e Employee2) DisplaySalary2() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, e.basicPay)
}

func(e Employee2) CalculatorLeavesLeft() int{
	return e.totalLeaves - e.leavesTaken
}


func TryInterface2(){
	em1 := Employee2{
		firstName:"bon",
		lastName: "jovi",
		basicPay: 100,
		totalLeaves: 10,
		leavesTaken: 5,
	}
	var empOp1 EmployeeOperations = em1
	empOp1.DisplaySalary2()
	fmt.Println("\nLeaves left =", empOp1.CalculatorLeavesLeft())
}


func hello(done chan bool){
	fmt.Println("Hello world goroutine")
	time.Sleep(2 * time.Second) //this will block hello process
	fmt.Println("Hello go routine awake and going to write to done")
	done <- true
}

func digits(number int, channel chan int){
	for number != 0 {
		digit := number % 10
		channel <- digit
		number /= 10
	}
	close(channel)
}

func calcSquares(number int, sequreOp chan int){
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch{
		sum += digit * digit
	}
	sequreOp <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}


func numbers(){
	for i := 1; i <= 5; i++{
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabets(){
	for i := 'a'; i <= 'e'; i++{
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

//unidirectional data channel
func sendData(sendch chan <- int){
	sendch <- 10
}

func goRoutineProducer(chnl chan int){
	for i:=0; i < 10; i++{
		time.Sleep(400 * time.Millisecond)
		chnl <- i
	}
	close(chnl)
}

func TryGoConcurrency(){
	done := make(chan bool)
	go hello(done)
	<- done

	numberValue := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(numberValue, sqrch)
	go calcCubes(numberValue, cubech)
	squares, cubes := <- sqrch, <- cubech
	fmt.Println("final output", squares + cubes)
	fmt.Println("main function")

	//single channel concurrency
	sendch := make(chan int)
	go sendData(sendch)
	fmt.Println("value from channel : ", <-sendch)


	//get data from goroutine producer
	channel1 := make(chan int)
	go goRoutineProducer(channel1)
	for {
		v, ok := <- channel1
		if ok == false{
			break;
		}
		fmt.Println("Received", v,ok)
	}
}

func write(ch chan int){
	for i := 0; i < 5; i++{
		ch <- i
		fmt.Println("successfully wrote ", i, " to ch")
	}
	close(ch)
}

func process(i int, wg *sync.WaitGroup){
	fmt.Println("started Goroutine", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func TryGoBufferChannel(){
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<- ch)
	fmt.Println(<- ch)

	//buffered channel
	ch2 := make(chan int, 2)
	go write(ch2)
	time.Sleep(2 * time.Second)
	for v := range ch2{
		fmt.Println("read value ", v , " from ch")
		time.Sleep(2 * time.Second)
	}

	//try wait group
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++{
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

type Job struct{
	id int
	randomno int
}
type Result struct{
	job Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits2(number int) int{
	sum := 0
	no := number
	for no != 0{
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup){
	for job := range jobs{
		output := Result{ job, digits2(job.randomno)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int){
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++{
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(numOfJobs int){
	for i :=0; i < numOfJobs; i++{
		randomNo := rand.Intn(99)
		job := Job{id : i, randomno : randomNo}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool){
	for result := range results{
		fmt.Printf("job id %d input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func TryLearnWorkerPool(){
	startTime := time.Now()
	noOfJobs := 20
	go allocate(noOfJobs)

	done := make(chan bool)
	go result(done)

	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<- done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}


func server1(ch chan string){
	time.Sleep(3 * time.Second)
	ch <- "from server 1"
}

func server2(ch chan string){
	time.Sleep(3 * time.Second)
	ch <- "from Server 2"
}

func processHeavyData(ch chan string){
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func TrySelect(){
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
		case s1 := <- output1:
			fmt.Println(s1)
		case s2 := <- output2:
			fmt.Println(s2)
	}

	//try process heavy data
	ch := make(chan string)
	go processHeavyData(ch)
	for{
		time.Sleep(1000 * time.Millisecond)
		select{
			case v := <- ch:
				fmt.Println("received value : ", v)
				return
			default:
				fmt.Println("no value received")
		}
	}
}

var x = 0;
func increment(wg *sync.WaitGroup){
	x = x + 1
	wg.Done()
}

func TryMutex(){
	var wg sync.WaitGroup
	for i := 0; i < 1000 ; i++{
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("final value of  x", x)
}

func allPersonSalary(personSallaries map[string]int){
	fmt.Println("=========================")
	fmt.Println("total length of persons : ", len(personSallaries))
	for key, value := range personSallaries{
		fmt.Printf("person sallary[%s] = %d\n", key, value)
	}
}


