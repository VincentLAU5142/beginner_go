package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

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
	// numbers[0] = 60
	fmt.Printf("This is our array %v\n", numbers)

	fmt.Println("This is the last value", numbers[len(numbers)-1])
	//NOT USE THIS COMMONLY
	// numbersAtInit := [...]int{10, 20, 30, 40, 50}

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("this is our matrix: %v\n", matrix)

	//slice
	//allNumbers := numbers[:]
	//firstThree := numbers[0:3]

	fruits := []string{"apple", "banana", "strawberry"}
	fmt.Printf("these are my fruits %v\n", fruits)

	fruits = append(fruits, "kiwi")
	fmt.Printf("these are my fruits with kiwi %v\n", fruits)

	fruits = append(fruits, "mango", "pineapple")
	fmt.Printf("these are my fruits with more fruits %v\n", fruits)

	moreFruits := []string{"blueberries", "tomato"}
	fruits = append(fruits, moreFruits...)
	fmt.Printf("these are my fruits with more fruits %v\n", fruits)

	//range of array, the _ are the index

	for index, value := range numbers {
		fmt.Printf("index %d and value %d\n", index, value)
	}
	capitalCities := map[string]string{
		"USA":   "Washington D.C.",
		"India": "New Delhi",
		"UK":    "London",
	}
	capital, exists := capitalCities["Germany"]
	if exists {
		fmt.Println("this is the capital", capital)
	} else {
		fmt.Println("Does not exist")
	}

	delete(capitalCities, "UK")
	fmt.Printf("this is the new deleted map:  %v\n", capitalCities)

	//struct
	person := Person{Name: "John", Age: 25}
	fmt.Printf("This is out person %+v\n", person)

	employee := struct {
		name string
		id   int
	}{
		name: "alice",
		id:   123,
	}

	type Address struct {
		Street string
		City   string
	}
	type Contact struct {
		Name    string
		Address Address
		Phone   string
	}
	contact := Contact{
		Name: "Marc",
		Address: Address{
			Street: "123 Main Street",
			City:   "Anytown",
		},
	}
	fmt.Println("this is contact", contact)
	fmt.Println("this is employee", employee)

	fmt.Println("name before: ", person.Name)
	// modifyPersonName(&person)
	person.modifyPersonName("Melkey") //individual input separate from the func
	fmt.Println("name after: ", person.Name)

	//quick demo in pointers to explain the modifyPersonName func
	x := 20
	ptr := &x
	fmt.Printf("value of x:= %d and address of x %p\n", x, ptr)
	*ptr = 30
	fmt.Printf("value of new x:= %d and address of x %p\n", x, ptr)
}

func add(a int, b int) int {
	return a + b
}

func calculateSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}

func (p *Person) modifyPersonName(name string) {
	p.Name = name
	//only in this scope are "melkey" ,if not use the * as the pointer and &
	fmt.Println("inside scope: new name", p.Name)
}
