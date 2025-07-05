package main

import (
	"fmt"
)

func main() {
	//Variables

	var name string = "Vincent"
	fmt.Printf("This is my name %s\n", name)

	// age := 25
	// fmt.Printf("This is my age %d\n", age)

	var city string
	city = "Seattle"
	fmt.Printf("This is my city %s\n", city)

	var country, continent string = "USA", "North America"
	fmt.Printf("This is my country %s and this is my continent %s\n", country, continent)

	var (
		isEmployed bool   = true
		salary     int    = 50000
		position   string = "developer"
	)
	fmt.Printf("isemployed: %t\n this is my salay: %d\n and this is my position: %s\n", isEmployed, salary, position)

	//Zero Values
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool

	fmt.Printf("int: %d float: %f string: %s and bool %t\n", defaultInt, defaultFloat, defaultString, defaultBool)

	const pi = 3.14

	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
	)

	fmt.Printf("Monday: %d - Tuesday: %d - Wednesday: %d\n", Monday, Tuesday, Wednesday)

	const typedAge int = 25
	const untypedAge = 25

	const (
		Jan = iota + 1 //1
		Feb            //2
		Mar            //3
		Apr            //4
	)

	fmt.Printf("Jan - %d feb - %d mar - %d apr - %d\n", Jan, Feb, Mar, Apr)

	//call the function add
	result := add(3, 4)
	fmt.Println("this is the result", result)

	sum, product := calculateSumAndProduct(10, 10)
	fmt.Printf("this is sum: %d and this is product: %d\n", sum, product)

	age := 30
	if age >= 18 {
		fmt.Println("you are an adult")
	} else if age >= 13 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("you are a child")
	}

	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("start of the week")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Midweek")
	case "Friday":
		fmt.Println("TGIF")
	default:
		fmt.Println("its the weekend")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("This is i", i)
	}

	// while loop
	counter := 0
	for counter < 3 {
		fmt.Println("this is the counter", counter)
		counter++
	}
	// iterations := 0
	// //infinitive loop
	// for {
	// 	if iterations > 3 {
	// 		break
	// 	}
	// 	iterations++
	// }

	// Arrays and Slices
	numbers := [5]int{10, 20, 30, 40, 50}
	numbers[0] = 60
	fmt.Printf("This is our array %v\n", numbers)

	fmt.Println("This is the last value", numbers[len(numbers)-1])
	//NOT USE THIS COMMONLY
	// numbersAtInit := [...]int{10, 20, 30, 40, 50}

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("this is our matrix: %v\n", matrix)
}

func add(a int, b int) int {
	return a + b
}

func calculateSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}
