package main

import (
	"fmt"
	"math"
	"strings"
)

// declaring struct
type Employee struct {
	name     string
	age      int
	division string
}

// Struct for Embedded Struct
type Person struct {
	name string
	age  int
}

type Employee2 struct {
	division string
	Person   Person
}

func main() {
	//types of declaring data types and giving its values
	var nama string = "Airell"
	var age int = 23

	fmt.Println("Ini adalah namanya ==>", nama)
	fmt.Println("Ini adalah umurnya ==>", age)

	var address string // you can declare with this type of method
	var email string

	address = "Britania"
	email = "Airell@gmai.com"

	fmt.Println("Ini adalah alamatnya ==>", address)
	fmt.Println("Ini adalah emailnya ==>", email)

	isCitizen := true // short declaration
	isCitizenID := "2345"

	fmt.Println("isCitizen? ==>", isCitizen)
	fmt.Println("Citizen ID ==>", isCitizenID)

	// multiple variable and its declaration
	var student1, student2, student3 string = "student1", "student2", "student3"

	var first, second, third int
	first, second, third = 1, 2, 3

	fmt.Println(student1, student2, student3)
	fmt.Println(first, second, third)

	// multiple variables but declaring different data types
	var student4, student4_age, isStudent4 = "student4", 23, true

	student5, student5_age, isStudent5 := "student5", 33, false

	fmt.Println(student4, student4_age, isStudent4)
	fmt.Println(student5, student5_age, isStudent5)

	// underscore variables (to prevent errors of unused variables)
	var firstVariable string

	var student6, student6_age, isStudent6 = "student6", 13, true

	_, _, _, _ = firstVariable, student6, student6_age, isStudent6

	//Printf usage: to identify data types
	student7, student8 := "student7", "student8"

	fmt.Printf("Tipe data variable first adalah %T \n", student7)
	fmt.Printf("Tipe data variable second adalah %T \n", student8)

	//Printf with more than one verb
	student9, student9_age, student9_address := "student9", 22, "Brimingham"

	fmt.Printf("Halo namaku %s, umurku adalah %d, dan aku tinggal di %s", student9, student9_age, student9_address)

	//constant data type: data type where its value cannot be modified
	const student10 string = "student10"
	const student10_age int = 10

	fmt.Println()
	fmt.Printf("name %s || age %d", student10, student10_age)
	fmt.Println()

	//math operators
	var value1 = 2 + 2
	var value2 = 2*3 + 2
	var value3 = (2 + 2) / 2
	var value4 = 6 % 2
	var value5 = 0
	value5++
	var value6 = 3
	value6--

	fmt.Println("value 1 =", value1)
	fmt.Println("value 2 =", value2)
	fmt.Println("value 3 =", value3)
	fmt.Println("value 4 =", value4)

	//relational operators
	var condition1 bool = 2 < 3
	var condition2 bool = "joey" == "Joey" // false due to caps sensitive
	var condition3 bool = 10 != 2.3
	var condition4 bool = 11 <= 13

	fmt.Println("first condition:", condition1)
	fmt.Println("second condition:", condition2)
	fmt.Println("third condition:", condition3)
	fmt.Println("fourth condition:", condition4)

	//logical operators
	var right = true
	var wrong = false

	var condition5 = wrong && right
	fmt.Printf("wrong and right \t(%t) \n", condition5)

	var condition6 = wrong || right
	fmt.Printf("wrong and right \t(%t) \n", condition6)

	var condition7 = !wrong
	fmt.Printf("wrong and right \t(%t) \n", condition7)

	//Array: fixed-sized sequences of elements
	//An array is a numbered sequence of elements of the same type with a fixed length
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var strings = [3]string{"nama1", "nama2", "nama3"}

	fmt.Println()
	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", strings)

	//Modifying Array element via index
	var fruits = [3]string{"apel", "pisang", "mangga"}
	fruits[0] = "Apple"
	fruits[1] = "banana"
	fruits[2] = "mango"

	fmt.Println()
	fmt.Printf("%#v\n", fruits)

	//Array looping elements
	var fruits1 = [3]string{"apple, banana", "mango"}

	for i, v := range fruits1 { // range lets you iterate through the "current" dimension
		fmt.Printf("Index : %d, Value: %s\n", i, v)
	}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index : %d, Value: %s\n", i, fruits1[i])
	}

	//Multidimensional Array
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

	//Slice: flexibel, dynamic views of arrays
	//Slices are flexible, dynamic sequences as you add or remove elements
	var fruits2 = []string{"apple", "banana", "mango"}

	_ = fruits2

	fmt.Printf("%#v", fruits2)

	fmt.Println()

	//Slice with make() function
	var fruits3 = make([]string, 3)

	_ = fruits3

	fmt.Printf("%#v", fruits3)

	fmt.Println()

	//Slice with append() function: to add values after in that dimension
	var fruits4 = make([]string, 3)

	fruits4 = append(fruits4, "apple", "banana", "manga")

	fmt.Printf("%#v", fruits4)

	fmt.Println()

	//Slice with append() with ellipsis(...): the three dot in append. used to add array to another array (one dimensional)
	var fruits5 = []string{"apple", "banana", "mango"}

	var fruits6 = []string{"durian", "pineapple", "starfruit"}

	fruits5 = append(fruits5, fruits6...)

	fmt.Printf("%#v", fruits5)

	fmt.Println()

	//Slice with copy() function
	var fruits7 = []string{"apple", "banana", "mango"}

	var fruits8 = []string{"durian", "pineapple", "starfruit"}

	fmt.Println("fruits7 =>", fruits7)

	nn := copy(fruits7, fruits8)

	fmt.Println("fruits7 =>", fruits7)
	fmt.Println("fruits8 =>", fruits8)
	fmt.Println("Copied elements =>", nn)

	//Slicing: to retrieve a certain index of a slice from certain element to certain element
	var fruits9 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var fruits9_1 = fruits9[1:4]
	fmt.Printf("%#v\n", fruits9_1)

	var fruits9_2 = fruits9[0:]
	fmt.Printf("%#v\n", fruits9_2)

	var fruits9_3 = fruits9[:3]
	fmt.Printf("%#v\n", fruits9_3)

	var fruits9_4 = fruits9[:] // is equals to fruits9[:len(fruits9)]
	fmt.Printf("%#v\n", fruits9_4)

	//Slicing and append combined
	var fruits10 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	fruits10 = append(fruits10[:3], "rambutan")

	fmt.Printf("%#v\n", fruits10)

	//Slice backing Array
	var fruits11 = []string{"apple", "mango", "durian", "banana", "starfruit"}

	var fruits12 = fruits11[2:4]

	fruits12[0] = "rambutan"

	fmt.Println("fruits11 => ", fruits11)
	fmt.Println("fruits12 => ", fruits12)

	//Slicing: cap function
	var fruits13 = []string{"apple", "mango", "durian", "banana"}

	fmt.Println("Fruits 13 cap:", cap(fruits13))
	fmt.Println("Fruits 13 len:", len(fruits13))

	var fruits14 = fruits13[0:3]

	fmt.Println("Fruits14 cap:", cap(fruits14))
	fmt.Println("Fruits14 len:", len(fruits14))

	var fruits15 = fruits13[1:]

	fmt.Println("Fruits14 len:", len(fruits15))
	fmt.Println("Fruits14 len:", len(fruits15))

	//Slice: creating a new backing array
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println("cars", cars)
	fmt.Println("newCars:", newCars)

	//Conditions: Temporary Variable
	var currentYear = 2021

	if age := currentYear - 1998; age < 17 { //how to declare variable inside conditions (variable used for temporary purposes)
		fmt.Println("Usia :", age)
		fmt.Println("Kamu belum boleh membuat kartu sim")
	} else {
		fmt.Println("Usia :", age)
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	}

	fmt.Println()

	//Switch
	var score = 8

	switch score {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("not bad")
	}

	fmt.Println()

	//Switch with relational operators
	score2 := 6

	switch { //notice switch declaration doesn't include variable due to various conditions
	case score2 == 8:
		fmt.Println("perfect")
	case (score2 < 8) && (score2 > 3):
		fmt.Println("not bad")
	default: //notice this to write more than one line
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}

	fmt.Println()

	//Switch with fallthrough keyword
	score3 := 6

	switch {
	case score3 == 8:
		fmt.Println("Perfect")
	case (score3 < 3) && (score > 3):
		fmt.Println("not bad")
		fallthrough
	case score3 < 5:
		fmt.Println("It is ok, but please study harder")
	default:
		{
			fmt.Println("Study harder")
			fmt.Println("You don't have a good score yet")
		}
	}

	fmt.Println()

	//Nested Conditions
	score4 := 10

	if score4 > 7 {
		switch score4 {
		case 10:
			fmt.Println("Perfect")
		default:
			fmt.Println("nice!")
		}
	} else {
		if score4 == 5 {
			fmt.Println("not bad")
		} else if score4 == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if score4 == 0 {
				fmt.Println("try harder!")
			}
		}
	}

	fmt.Println()

	//Looping: first method (for loop)
	for i := 0; i < 3; i++ {
		fmt.Println("angka", i)
	}

	fmt.Println()

	//Looping: second method (basic for loop)
	var i = 0

	for i < 3 {
		fmt.Println("Angka", i)
		i++
	}

	fmt.Println()

	//Looping: third method (stopping condition using break)
	var i2 = 0

	for {
		fmt.Println("Angka", i2)
		i2++

		if i2 == 3 {
			break
		}
	}

	fmt.Println()

	//Looping: break and continue
	for i := 1; i <= 10; i++ {
		if i%2 == 1 { //condition if i is odd number
			continue //tells you to keep counting and due to skip println
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	fmt.Println()

	//Nested Looping
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	fmt.Println()

	//Looping: Label
outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke - ", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}

	//Function
	greet("Airell", "Jalan Sudirman")

	// calling functions: without return
	greet("Airell", "jalan Sudirman")
	fmt.Println()

	// calling functions: with return
	var names = []string{"Airell", "Jordan"}
	var printMsg = greet2("Heii", names)

	fmt.Println(printMsg)
	fmt.Println()

	// calling functions: for multiple return
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Println("Area: ", area)
	fmt.Println("Circumference: ", circumference)
	fmt.Println()

	// calling functions: for predefined values
	var diameter2 float64 = 15
	var area2, circumference2 = calculate2(diameter2)

	fmt.Println("Area: ", area2)
	fmt.Println("Circumference: ", circumference2)
	fmt.Println()

	//calling variadic functions with map method
	studentLists := print("Airell", "Nanda", "Mailo", "Schannel", "Marco")

	fmt.Printf("%v", studentLists)
	fmt.Println()

	//calling variadic function with array method
	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}

	result := sum(numberLists...)

	fmt.Println("Results:", result)
	fmt.Println()

	//calling variadic function with struct
	profile("Airell", "pasta", "ayam geprek", "ikan roa", "sate padang")

	//Pointers: Memory Address
	var firstNumber int = 4
	var secondNumber *int = &firstNumber

	fmt.Println("firstNumber (value) :", firstNumber)
	fmt.Println("firstNumber (memori address) :", &firstNumber)

	fmt.Println("secondNumber (value) :", *secondNumber)
	fmt.Println("secondNumber (memori address) :", secondNumber)

	fmt.Println()

	//changing value through pointer
	var firstPerson string = "John"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("secondNumber (value) :", *secondPerson)
	fmt.Println("secondNumber (memori address) :", secondPerson)

	*secondPerson = "Doe"

	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("secondNumber (value) :", *secondPerson)
	fmt.Println("secondNumber (memori address) :", secondPerson)

	fmt.Println()

	//using pointer as function parameter
	var a int = 10

	fmt.Println("Before:", a)

	changeValue(&a)

	fmt.Println("After:", a)

	fmt.Println()

	//Struct: Collectible data type
	var employee Employee

	employee.name = "Airell"

	employee.age = 23

	employee.division = "Curriculum Developer"

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)
	fmt.Println()

	//Initializing struct: method 1
	var employee1 = Employee{}
	employee1.name = "Airell"
	employee1.age = 23
	employee1.division = "Curriculum Developer"

	fmt.Printf("Employee1: %+v\n", employee1)

	//Initializing struct: method 2
	var employee2 = Employee{name: "Ananda", age: 23, division: "Finance"}

	fmt.Printf("Employee2: %+v\n", employee2)

	fmt.Println()

	//Pointer to a struct
	var employee3 = Employee{name: "Airell", age: 23, division: "Curriculum Developer"}

	var employee4 *Employee = &employee3

	fmt.Println("Employee3 name:", employee3.name)
	fmt.Println("Employee4 name:", employee4.name)

	employee4.name = "Ananda"

	fmt.Println("Employee3 name:", employee3.name)
	fmt.Println("Employee4 name:", employee4.name)

	fmt.Println()

	//Embedded Struct
	var employee5 = Employee2{}

	employee5.Person.name = "Airell"
	employee5.Person.age = 23
	employee5.division = "Curriculum Developer"

	fmt.Printf("%+v", employee5)

	fmt.Println()

	//Anonymous Struct without field insertion
	var employee6 = struct {
		person   Person
		division string
	}{}

	employee6.person = Person{name: "Airell", age: 23}
	employee6.division = "Curriculum Developer"

	fmt.Printf("Employee6: %+v\n", employee6)
	fmt.Println()

	//Anonymous struct with field insertion
	var employee7 = struct {
		person   Person
		division string
	}{
		person:   Person{name: "Ananda", age: 23},
		division: "Finance",
	}

	fmt.Printf("Employee7: %+v\n", employee7)
	fmt.Println()

	//Slice of struct
	var people = []Person{
		{name: "Airell", age: 23},
		{name: "Ananda", age: 23},
		{name: "Mailo", age: 23},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}

	fmt.Println()

	//combining Anonymous struct with Slice
	var employee8 = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "Airell", age: 23}, division: "Curriculum Developer"},
		{person: Person{name: "Ananda", age: 23}, division: "FInance"},
		{person: Person{name: "Mailo", age: 23}, division: "Marketing"},
	}

	for _, v := range employee8 {
		fmt.Printf("%+v\n", v)
	}

	fmt.Println()

	//calling method
	var person = Person{name: "Airell", age: 23}

	fmt.Println(person.Introduce("Hello everyone!!"))

	fmt.Println()

	//calling method with Pointers
	var person1 = Person{name: "Airell", age: 23}

	person1.ChangeName1()
	fmt.Println("Change name with ChangeName1 method", person1.name)

	person1.ChangeName2()
	fmt.Println("Change name with ChangeName2 method", person1.name)

	fmt.Println()

	//calling Pointer Method
	var person3 = Person{name: "Airell", age: 23}

	person3.Greet()

	var person4 = &Person{name: "Mailo", age: 15}
	person4.Greet()

	//Maps: key-value stores implemented hash tables
	//Map is Go's built-in associative data type to stores key-value pairs with fast average-time lookups
	//Maps let you use more meaningful keys such as names, IDs, or other comparable values
}

// declaring functions without return
func greet(name, address string) {
	fmt.Println("Hello there: My name is", name)
	fmt.Println("I live in", address)
}

// declaring functions with return
func greet2(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")

	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result
}

// declaring functions with multiple return
func calculate(d float64) (float64, float64) {
	var area float64 = math.Pi * math.Pow(d/2, 2)
	var circumference = math.Pi * d
	return area, circumference
}

// declaring functions with predefined return value
func calculate2(d float64) (area2 float64, circumference2 float64) {
	area2 = math.Pi * math.Pow(d/2, 2)
	circumference2 = math.Pi * d
	return
}

// declaring variadic function that collects results into a slice of maps
func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}

	return result
}

// declaring variadic function by unpacking a slice with ellipsis (...)
func sum(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}
	return total
}

// declaring variadic function by passing individual arguments directly
func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ",")

	fmt.Println("Hello there!!! I'm ", name)
	fmt.Println("I really love to eat ", mergeFavFoods)
}

func changeValue(number *int) {
	*number = 20
}

// declaring Method: function that is attached with data type
func (p Person) Introduce(msg string) string {
	return fmt.Sprintf("%s My name is %s and I'm %d years old", msg, p.name, p.age)
}

// declaring method with Pointers
func (p Person) ChangeName1() {
	p.name = "Mailo"
}

func (p *Person) ChangeName2() {
	p.name = "Mailo"
}

// declaring Pointer method
func (p *Person) Greet() {
	fmt.Println("Haii everyone", p)
}
