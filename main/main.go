package main

//looks like python but legit
import (
	"fmt"
	"math/rand"
	"math"
	"runtime"
	"time"
	"strings"
	"golang.org/x/tour/wc"
)

//this is function , short way of (x int, y int)
func add(x , y int) int{
	return x + y
}

func swap(x, y string) (string, string){
	return y, x
}

//naked return will be harmful for reading
func split(sum int) (x, y int){
	x = sum * 4/ 9;	// 4/9 ratio
	y = sum - x;
	return
}

//variables in the packages, this has default value
var c, python, java bool;

func main(){
	fmt.Println("Hello, 世界")
	fmt.Println("my favourite number is ",rand.Intn(10))
	fmt.Printf("Now you have %g problems \n", math.Sqrt(7))
	fmt.Println(math.Pi)

	fmt.Println("total 10 + 30 : ", add(10, 30));

	a, b := swap("hello", "world");
	fmt.Println(a, b);

	fmt.Println(split(100));


	//this has default value
	var i , ce , e = 10, true, "bool";
	ce2 := 10;
	fmt.Println(i,c, java, python, ce, e);
	fmt.Println("ce2 :", ce2);

	var x, y int = 3, 4;
	var f float64 = math.Sqrt(float64(x * x + y * y))
	var z uint = uint(f);
	fmt.Println(x,y,z);

	learnConstants();
	simpleLogic();
	simplePointers();
	simpleMoreTypes();
	wc.Test(WordCount);

	function_compute_main();
}

func function_compute(fn func(float64, float64)float64) float64{
	return fn(3, 4);
}

func function_compute_main(){
	hypot := func(x, y float64) float64{
		return math.Sqrt(x*x + y*y);
	}

	fmt.Println(hypot(5, 12));
	fmt.Println("compute data : ", function_compute(hypot));

	//HERE :  https://tour.golang.org/moretypes/25
}




type Vertex struct{
	x int
	y int
}

func WordCount(s string)map[string]int{
	return map[string]int{"x" : 1};
}

func simpleMoreTypes(){
	v := Vertex{1, 2};
	v1 := Vertex{x : 1};
	fmt.Println("v1 : ", v1.x);
	fmt.Println("v1 : ", v1.y);

	fmt.Println("x : ", v.x);
	fmt.Println("y : ", v.y);

	var a [2]string;
	a[0] = "Hello";
	a[1] = "world";
	fmt.Println(a[0], a[1]);
	fmt.Println(a);

	primes := [6]int{2,3,5,7, 11, 13};
	fmt.Println("primes : ", primes);

	//slice is by references
	var s[]int = primes[0:3];
	fmt.Println("s value : ",s);

	s[0] = 100;
	fmt.Println("s[0] value : ", s[0], primes[0]);

	//slice literal
	q := []int{2,3,4,5,7, 11};
	r := []bool{true, false, true, true, false, true};
	fmt.Println(q,r);

	aStruct := []struct{
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, false},
	}
	fmt.Println("aStruct : ", aStruct);

	q1 := q[2:];
	fmt.Println("q : ",q1);
	fmt.Println("q : ", q[:]);
	fmt.Println("q : ", q[:3]);
	printSlices(q);
	printSlices(q1);
	nilSlices();
	makeSlices();
	boardOfSlice();
	appendSlice();
	rangeData();
}

func rangeData(){
	//range
	var pow = []int{1,2,4,8,16,32,64,128};
	for i,v := range pow{
		fmt.Printf("2**%d = %d\n",i, v);
	}

	for _, value := range pow{
		fmt.Printf("value of pow : %d\n", value);
	}


	//struct
	type Vertex struct{
		Lat, Long float64
	};

	var m map[string]Vertex;
	m = make(map[string]Vertex);
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	};
	fmt.Println("Bell Labs : ", m["Bell Labs"]);
	fmt.Println("============");

	var mapLiteral = map[string]Vertex{
		"Bell Labs": Vertex{40.68433, -74.39967 },
		"Google" : Vertex{37.42202, -122.08488},
		"Testing" : Vertex{Lat: 37.1034, Long: 233.111},
	};
	fmt.Println("map literal : ", mapLiteral);

	m1 := make(map[string]int);
	m1["answer"] = 42;
	m1["answer"] = 48;
	fmt.Println("answer : ", m1["answer"]);


	v , ok := m["answer"];
	fmt.Println("answer : ", v, ok);
}

func appendSlice(){
	var s[]int;
	printSlices(s);

	s = append(s, 0);
	printSlices(s);

	s = append(s, 2,3,4);
	printSlices(s);
}

func makeSlices(){
	a := make([]int, 5);
	printSlices(a);
	fmt.Println("a : ", a[0:cap(a)]);
}

func boardOfSlice(){
	//slice 2nd dimention array
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	};
	board[0][0] = "X";
	board[2][2] = "O";
	board[1][2] = "X";

	fmt.Println("board :");
	fmt.Println("=========");
	for i:= 0; i < len(board); i++{
		fmt.Printf("%s\n", strings.Join(board[i], " "));
	}
}

func nilSlices(){
	var s[] int;
	fmt.Println(s, len(s), cap(s));
	if s == nil{
		fmt.Println("nil!!")
	}
}


func printSlices(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s);
}

//lesson variables
func learnConstants(){
	const world = "世界";
	fmt.Println("Happy", world, "Day");
}


func simpleLogic(){
	sum := 0;
	for i :=0 ; i < 10; i++{
		sum += 1;
	}
	fmt.Println(sum);

	//for without condition
	i2 := 0
	for i2 < 10{
		i2+= 1;
	}
	fmt.Println("i2 :", i2);


	fmt.Println("sqrt value : ",sqrt(2));
	fmt.Println("sqrt value : ", sqrt(-4));
	fmt.Println("pow value : ", pow(3,2 , 10));
	fmt.Println("pow value : ", pow(3,3 , 20));
	selectOS();
}

func sqrt(x float64) string{
	if x < 0{
		return sqrt(-x) + "i";
	}

	return fmt.Sprint(math.Sqrt(x));
}

func pow(x, n, lim float64) float64{
	if v := math.Pow(x, n); v < lim{
		return v;
	}
	return lim;
}



//runtime stuff
func selectOS(){
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin" :
		fmt.Println("OS X");
	case "linux" :
		fmt.Println("Linux");
	default:
		fmt.Println("%s", os);
	}

	dayEvaluationOrder();
}


func dayEvaluationOrder(){
	fmt.Println("When's Saturday ?");
	today := time.Now().Weekday();
	switch time.Saturday {
	case today + 0:
		fmt.Println("today.");
	case today + 1:
		fmt.Println("Tomorrow.");
	case today + 2:
		fmt.Println("In two days");
	default:
		fmt.Println("Too Far Away");
	}

	t := time.Now();
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning !");
	case t.Hour() < 17:
		fmt.Println("Good Afternoon !");
	default:
		fmt.Println("Good Evening.");
	}

	deferFunctions();
}


func deferFunctions(){
	fmt.Println(", ");
	defer fmt.Println("world");
	fmt.Println("hello");

	fmt.Println("counting...");
	for i:= 0; i < 10; i++{
		defer fmt.Println(i);
	}
	fmt.Println("done")
}



func simplePointers(){
	i, j := 42, 270;

	p := &i;
	fmt.Println(*p);
	*p = 2;
	fmt.Println(i);
	fmt.Println(j);
}

