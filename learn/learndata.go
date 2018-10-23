package learn

import "fmt"

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

func tryInterface2(){

}


func allPersonSalary(personSallaries map[string]int){
	fmt.Println("=========================")
	fmt.Println("total length of persons : ", len(personSallaries))
	for key, value := range personSallaries{
		fmt.Printf("person sallary[%s] = %d\n", key, value)
	}
}


