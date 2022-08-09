package main

//this is the package scope
import (
	"fmt"
	"reflect"
	"strconv"
)

var l int = 50

var (
	actorName string = "Ado no"
	companion string = "aiyo manda"
	quote     string = "phone eka lamayata diyan"
	version   int    = 69
)

var m int = 10

// enum const
const (
	q = iota
	w = iota
	e = iota
)

// struct
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

// composition
type Animal struct {
	Name   string `required max: "100"`
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {

	var j int = 58 //default int is 42 bits
	//var p uint = 42 //unsigned integer value range starts from 0 not -x (?)

	var k int
	k = 100

	i := 42

	var m int = 5 //shadowing

	var n string = strconv.Itoa(i)

	fmt.Printf("%v , %T \n", n, n)

	bb := 1 == 1
	cc := 1 == 2

	var dd bool //uninitialized bool var is automatically false. numerics will be 0

	fmt.Printf("%v , %T \n", bb, bb)
	fmt.Printf("%v , %T \n", cc, cc)
	fmt.Printf("%v , %T \n", dd, dd)

	fmt.Println(i + j + k + l + m)

	//bit operators

	a := 10 // 1010
	b := 3  // 0011

	fmt.Println(a & b)  // and a and b binary 0010 = 2
	fmt.Println(a | b)  // a or b binary 1011 = 11
	fmt.Println(a ^ b)  //exclusive or 1001 (if both have same it gets removed) = 9
	fmt.Println(a &^ b) //and not | a & (^ b) so 1010 and 1100 = 1000

	// bit shift

	c := 8 // 2 ^ 3

	fmt.Println(c << 3) // 2^3 * 2^3 = 2^6 = 64
	fmt.Println(c >> 3) // 2^3 / 2^3 = @^0

	//float
	oo := 3.14
	oo = 13.7e72
	oo = 2.1e14
	fmt.Printf("%v, %T\n", oo, oo)

	//complex type
	var im complex64 = complex(5, 12) // or 5 + 12i
	fmt.Printf("%v, %T\n", im, im)

	//string
	str := "this is a string"
	sb := []byte(str)
	fmt.Printf("%v, %T\n", sb, sb)

	//constant
	const myConst int = 100
	fmt.Printf("%v, %T\n", myConst, myConst)

	//Arrays and Slices
	grade := [...]int{97, 85, 83} //array - they dont have pointers when it's assigned to another var
	fmt.Printf("Grade: %v\n", grade)
	grade[1] = 100
	fmt.Printf("Grade: %v\n", grade)
	gradeNew := grade
	gradeNew[1] = 222
	fmt.Printf("Grade: %v\n", gradeNew)

	brade := []int{97, 85, 83} //slice - they have pointers when it's assigned to another var
	fmt.Printf("Grade: %v\n", brade)
	brade[1] = 100
	fmt.Printf("Grade: %v\n", brade)
	bradeNew := brade
	bradeNew[1] = 222
	fmt.Printf("Grade: %v\n", bradeNew)

	//Maps and Structs
	kvMap := make(map[string]int)
	kvMap = map[string]int{
		"Key_1": 1,
		"Key_2": 2,
		"Key_3": 3,
		"Key_4": 4,
	}

	kvMap["Key_5"] = 5
	delete(kvMap, "Key_4")

	_, ok := kvMap["Key_4"] // wht this does is check the map for key and if it doesnt exist return false

	fmt.Println(kvMap)
	fmt.Println(ok)

	//stuct - these refer to independed instances (transfer to new var mean new instance)
	aDoctor := Doctor{
		number:    3,
		actorName: "Ane Manda John",
		companions: []string{
			"weeeeee",
			"woooooo",
		},
	}

	fmt.Println(aDoctor.actorName)

	//Composition
	//check package level

	// bababui := Bird{}
	// bababui.Name = "Emu"
	// bababui.Origin = "Australia"
	// bababui.SpeedKPH = 48
	// bababui.CanFly = false

	bababui := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}

	fmt.Println(bababui)

	//tags
	//check animal struct
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

	//if
	number := 50
	guess := 30

	if guess < number {
		fmt.Println("Too low")
	} else if guess > number {
		fmt.Println("Too high")
	} else if guess == number {
		fmt.Println("Correct!")
	}

	//switch
	switch i := 1 + 2; i {
	case 1, 3:
		fmt.Println("one or three")
	case 2, 4:
		fmt.Println("two or four")
	default:
		fmt.Println("not one or two")
	}

	hek := 10
	switch {
	case hek <= 10:
		fmt.Println("one or three")
		fallthrough // this will go to next case ignoring logic
	case hek <= 20:
		fmt.Println("one or three")
	default:
		fmt.Println("one or three")
	}

	var iter interface{} = 1
	switch iter.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("another type")
	}

	//for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//for loop
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

Loop: //label
	for i := 1; i <= 5; i++ {
		for j := 1; i <= 5; j++ {
			fmt.Println(i * j)
			if i*j >= 5 {
				break Loop //label to break to if not this will go to outter loop and continue
			}
		}
	}

	//for range loop for slices
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	//defer panic recovery
	fmt.Println("start")
	defer fmt.Println("middle") // moves this fuctions execution after the main function not to bottom
	fmt.Println("end")

	//defer behaves in last in first out

	//pointers
	var point int = 42
	var boint *int = &a // * = pointer here & = address of
	fmt.Println(point, boint)
	a = 12
	fmt.Println(point, *boint) // * = is a dereferce here

	//pointer arithmatic
	apoint := [3]int{1, 2, 3}
	bpoint := &apoint[0]
	cpoint := &apoint[1]
	fmt.Printf("%v %p %p\n", apoint, bpoint, cpoint)

	var ms *myStruct
	ms = new(myStruct)
	ms.foo = 42
	fmt.Println(ms.foo)

	//functions
	greeting := "Hello"
	name := "Stacey"
	sayGreeting(&greeting, &name)
	fmt.Println(name)

	result := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(result)

	result1, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result1)

	//anon func
	func() {
		fmt.Println("Anon Func!")
	}()

	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()

	//Interfaces
	var wr Writer = ConsoleWriter{}
	wr.Write([]byte("Hello GO!"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// this is for pointer example
type myStruct struct {
	foo int
}

//NOTE: Maps and Slices vars have addresses for thier underlying arrays
//so those vars will behave like pointer by default. be careful!

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

func sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("Sum: ", result)

	return result
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

// this is method (as in a methods in objects.. bit different from func)
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
