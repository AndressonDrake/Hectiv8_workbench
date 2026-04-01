package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"testing_go/helpers"
	"time"
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

// declaring Interface
type shape interface {
	area() float64
	perimeter() float64
}

// declaring struct for rectangle properties
type rectangle struct {
	width, height float64
}

// declaring struct for circle properties
type circle struct {
	radius float64
}

type Notification struct {
	UserID  int
	Message string
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
	fmt.Printf("wrong or right \t(%t) \n", condition6)

	var condition7 = !wrong
	fmt.Printf("not wrong \t(%t) \n", condition7)

	//Array: fixed-sized sequences of elements
	//An array is a numbered sequence of elements of the same type with a fixed length
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var strArr = [3]string{"nama1", "nama2", "nama3"}

	fmt.Println()
	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", strArr)

	//Modifying Array element via index
	var fruits = [3]string{"apel", "pisang", "mangga"}
	fruits[0] = "Apple"
	fruits[1] = "banana"
	fruits[2] = "mango"

	fmt.Println()
	fmt.Printf("%#v\n", fruits)

	//Array looping elements
	var fruits1 = [3]string{"apple", "banana", "mango"}

	for i, v := range fruits1 { // range lets you iterate through the "current" dimension
		fmt.Printf("Index : %d, Value: %s\n", i, v)
	}

	for i := 0; i < len(fruits1); i++ {
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

	fmt.Println("Fruits14 cap:", len(fruits15))
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
	case (score3 < 3) && (score3 > 3):
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

	fmt.Println()

	// calling exported functions from herlpers directory
	helpers.GreetHelpers() // imports/exports functions only work with first letter capital

	fmt.Println()

	// calling interface shape
	var c1 shape = circle{radius: 5}
	var r1 shape = rectangle{width: 3, height: 2}

	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of r1: %T\n", r1)

	fmt.Println()

	fmt.Println("Circle area", c1.area())
	fmt.Println("Circle Perimeter", c1.perimeter())

	fmt.Println("Rectangle area", r1.area())
	fmt.Println("Rectangle perimeter", r1.perimeter())

	fmt.Println()

	// declaring interface with type assertion
	var c2 shape = circle{radius: 5}

	value, ok := c2.(circle)

	if ok == true {
		fmt.Printf("Circle value: %T\n", value)
		fmt.Printf("Circle volume: %v\n", value.volume())
	}

	// Empty Interface
	var randomValues interface{}

	_ = randomValues

	randomValues = "Jalan Sudirman"

	randomValues = 20

	randomValues = true

	randomValues = []string{"Airell", "Nanda"}

	fmt.Printf("randomValues : %T\n", randomValues)
	fmt.Printf("randomValues : %t\n", randomValues)
	fmt.Printf("randomValues : %d\n", randomValues)
	fmt.Printf("randomValues : %s\n", randomValues)
	fmt.Printf("randomValues : %v\n", randomValues)

	fmt.Println()

	// Empty Interface with Type Assertion
	var v interface{}

	v = 20

	if value, ok := v.(int); ok == true {
		v = value * 9
	}

	fmt.Printf("v : %d\n", v)

	fmt.Println()

	//Empty Interface with map & slice
	rs := []interface{}{1, "Airell", true, 2, "Ananda", true}

	rm := map[string]interface{}{
		"Name":   "Airell",
		"Status": true,
		"Age":    23,
	}

	_, _ = rs, rm

	fmt.Printf("rs : %v\n", rs)
	fmt.Printf("rm : %v\n", rm)

	fmt.Println()

	//Concurrency
	numbersConcur := []int{5, 7, 3, 10}

	var wg sync.WaitGroup
	wg.Add(len(numbersConcur))

	for _, num := range numbers {
		go factorial(num, &wg)
	}

	wg.Wait()
	fmt.Println("All factorials calculated")

	fmt.Println()

	//Goroutines
	fmt.Println("main execution started")

	go firstProcess(8)

	secondProcess(8)

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	fmt.Println("main execution ended")

	fmt.Println()

	//Gorountine Asynchronus process
	fmt.Println("main execution started")

	go firstProcess1(8)

	secondProcess1(8)

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	time.Sleep(time.Second * 2)

	fmt.Println("main execution ended")

	fmt.Println()

	//Runtime differenes between synchronus and asynchronus
	now := time.Now()

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("HelloWorld")
	}

	diff := time.Since(now)
	fmt.Println(diff)

	fmt.Println()

	wg1 := &sync.WaitGroup{}
	now1 := time.Now()

	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println("HelloWorld")
			wg1.Done()
		}()
	}

	wg1.Wait()
	diff1 := time.Since(now1)
	fmt.Println(diff1)

	fmt.Println()

	//Concurrency example: Asynchronus Email Sending
	notifications := []Notification{
		{UserID: 101, Message: "Your Order has been confirmed"},
		{UserID: 202, Message: "Your account has been created"},
		{UserID: 303, Message: "Your payment was successful"},
	}

	for _, notification := range notifications {
		go sendEmailAsync(notification.UserID, notification.Message)
	}

	fmt.Println("Main application continues...")

	time.Sleep(3 * time.Second)

	fmt.Println("Main application finished")

	fmt.Println()

	//Scenario Image Processing Services
	images := []string{
		"https://example.com/image1.png",
		"https://example.com/image2.png",
		"https://example.com/image3.png",
		"https://example.com/image4.png",
	}

	for _, imageURL := range images {
		go processImage(imageURL)
	}

	fmt.Println("Image processing started, main application continues...")

	time.Sleep(5 * time.Second)

	fmt.Println("All image processing completed")

	//Scenario: Task Scheduler
	var wg2 sync.WaitGroup

	task1 := func() {
		fmt.Println("Task 1 is being executed")
	}
	task2 := func() {
		fmt.Println("Task 2 is being executed")
	}
	task3 := func() {
		fmt.Println("Task 3 is being executed")
	}

	wg2.Add(3)
	scheduleTask(task1)
	scheduleTask(task2)
	scheduleTask(task3)

	fmt.Println("Main application continues...")

	go func() {
		wg2.Wait()
		fmt.Println("All tasks completed")
	}()

	wg2.Done()
	wg2.Done()
	wg2.Done()

	//Scenario: Download Manager
	downloadJobs := map[string]string{
		"https://www.youtube.com/watch?v=a3ICNMQW7Ok&pp=ygUNZXhhbXBsZSB2aWRlbw%3D%3D": "video1.mp4",
		"https://www.youtube.com/watch?v=K4TOrB7at0Y&pp=ygUNZXhhbXBsZSB2aWRlbw%3D%3D": "video2.mp4",
		"https://www.youtube.com/watch?v=aGTg9xsI8oY&pp=ygUNZXhhbXBsZSB2aWRlbw%3D%3D": "video3.mp4",
	}

	var wg3 sync.WaitGroup
	wg3.Add(len(downloadJobs))

	for url, destination := range downloadJobs {
		go func(u, d string) {
			defer wg3.Done()
			downloadFile(u, d)
		}(url, destination)
	}

	wg3.Wait()

	fmt.Println("All files downloaded")

	//Channels: Implementing channels
	c := make(chan string)

	go introduce("Airell", c)

	go introduce("Nanda", c)

	go introduce("Mailo", c)

	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	msg3 := <-c
	fmt.Println(msg3)

	close(c)

	//Channels with anonymous function
	c3 := make(chan string)

	students := []string{"Airell", "Kailo", "Indah"}

	for _, v := range students {
		go func(student string) {
			fmt.Println("Student", student)
			result := fmt.Sprintf("Hai, my name is %s", student)
			c3 <- result
		}(v)
	}

	for i := 1; i <= 3; i++ {
		printChannel(c3)
	}

	close(c3)

	fmt.Println()

	//Channels with directions
	c4 := make(chan string)

	studentsChannel := []string{"Airell", "Kaito", "indah"}

	for _, v := range studentsChannel {
		go introduce1(v, c4)
	}

	for i := 1; i <= 3; i++ {
		printChannel1(c4)
	}

	close(c4)

	fmt.Println()

	//Channels: unbuffered channel
	c5 := make(chan int)

	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c5 <- 10
		fmt.Println("func goroutine after sending data into the channel")
	}(c5)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine starts receiving data")
	d := <-c5
	fmt.Println("main goroutine recieved data", d)

	close(c5)
	time.Sleep(time.Second)

	fmt.Println()

	//Channels: buffered channel
	c6 := make(chan int, 3)

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
			c6 <- i
			fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
		}

		close(c)
	}(c6)

	fmt.Println("main goroutine sleeps 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c6 {
		fmt.Println("main goroutine received value from channel:", v)
	}

	fmt.Println()

	//Channle: Select
	c7 := make(chan string)
	c8 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)

		c7 <- "Hello!"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c8 <- "Salut!"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-c7:
			fmt.Println("Received", msg1)
		case msg2 := <-c8:
			fmt.Println("Received", msg2)
		}
	}

	fmt.Println()

	//Scenario: Sum of Squares concurrent calculation
	numbersChannel := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	resultCh := make(chan int)

	mid := len(numbersChannel) / 2
	firsthalf := numbersChannel[:mid]
	secondHalf := numbersChannel[mid:]

	var wgCh sync.WaitGroup
	wgCh.Add(2)
	resultCh1 := 0

	go calculateSumOfSquares(firsthalf, resultCh, &wgCh)
	go calculateSumOfSquares(secondHalf, resultCh, &wgCh)

	go func() {
		wgCh.Wait()
		close(resultCh)
	}()

	for sum := range resultCh {
		resultCh1 += sum
	}

	fmt.Printf("Sum of squares : %d\n", resultCh1)

	//Scenario: Concurrent Word Count
	filenames := []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
	}

	var wg4 sync.WaitGroup
	wg4.Add(len(filenames))

	wordCounts := make(map[string]int)

	for _, filename := range filenames {
		go func(fn string) {
			defer wg4.Done()
			ch := make(chan int)
			go countWords(fn, ch)
			wordCounts[fn] = <-ch
		}(filename)
	}

	wg4.Wait()

	for filename, count := range wordCounts {
		fmt.Printf("%s: %d words\n", filename, count)
	}

	fmt.Println()

	//Defer
	defer fmt.Println("defer function starts to execute")
	fmt.Println("Hai everyone")
	fmt.Println("Welcome back to Go learning center")

	fmt.Println()

	//Defer using functions
	callDeferFunc()
	fmt.Println("Hai everyone !!")

	fmt.Println()

	//Scenario: Graceful Shutdown with Defer
	server := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("Server error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down the server...")
	if err := server.Shutdown(context.Background()); err != nil {
		fmt.Println("Error while shutting down the server:", err)
	}
	fmt.Println("Server gracefully stopped.")

	fmt.Println()

	//Scenario: Deferred File Closing
	filename := "data.txt"
	lines, err := processFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("NUmber of lines in %s: %d\n", filename, lines)

	fmt.Println()

	//Scenario: Defer for Cleanup Tasks
	performTask("Task 1")
	performTask("Task 2")
	performTask("Task 3")

	//Error data type
	var number int
	var err1 error

	number, err1 = strconv.Atoi("123GH")

	if err1 == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err1.Error())
	}

	number, err1 = strconv.Atoi("123")

	if err1 == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err1.Error())
	}

	//Declaring custom error and panic
	defer catchErr()

	var password string

	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		panic(err)
	} else {
		fmt.Println(valid)
	}

	//Recover function
	var password1 string

	fmt.Scanln(&password1)

	if valid, err := validPassword1(password1); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}

	//Scenario: Graceful Exit with Panic Recovery
	fmt.Println("5 / 2 =", divide(5, 2))
	fmt.Println("10 / 0 =", divide(10, 0))
	fmt.Println("8 / 2 =", divide(8, 2))

	//Exit
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before Exiting")
	os.Exit(1)

	fmt.Println()

	//Maps: key-value stores implemented hash tables
	//Map is Go's built-in associative data type to stores key-value pairs with fast average-time lookups
	//Maps let you use more meaningful keys such as names, IDs, or other comparable values
}

/*
=================================================================================
CONCURRENCY
=================================================================================

The ability to manage multiple tasks independently without a fixed order.
Tasks don't have to run at the exact same time — they just progress without
waiting on each other.

Concurrency    → organizing tasks to progress independently (dealing with many)
Asynchronism   → completing tasks without blocking, efficient for I/O-bound ops
Parallelism    → tasks literally run at the same time across multiple CPU cores

In Go, concurrency is built around two things:
  - Goroutines : lightweight threads launched with the `go` keyword
  - Channels   : pipeline for Goroutines to communicate and sync safely

=================================================================================
CONCURRENCY vs ASYNCHRONISM vs PARALLELISM
=================================================================================

CONCURRENCY
  - Focuses on organizing and managing tasks so they can progress independently.
  - Concurrent tasks CAN run at the same time, but not always.
  - Achieved using techniques like multithreading or Goroutines in Go.
  - Helps improve resource utilization and overall performance.

ASYNCHRONISM
  - Refers to a method or pattern for completing tasks without blocking.
  - Allows tasks to operate without waiting for one another to finish.
  - Often event-driven and highly efficient for I/O-bound operations (input/output).
  - In I/O scenarios, asynchronism is usually combined with concurrency
    to further improve resource performance.

PARALLELISM
  - Allows multiple tasks to run *truly simultaneously* across multiple
    processors or CPU cores.
  - Tasks run at the exact same time, not just independently progressing.
  - Works best for CPU-bound tasks that can be split into independent sub-tasks.
  - Combining multiple processing units via parallelism reduces overall
    execution time.

CONCURRENCY  → about *dealing with* lots of things at once
               (tasks take turns, managed independently, order is unknown)
PARALLELISM  → about *doing* lots of things at once
               (tasks literally run at the same time on separate cores)

=================================================================================
HOW CONCURRENCY WORKS
=================================================================================

A concurrent program breaks a large task into smaller sub-tasks.
These sub-tasks can be executed concurrently by multiple threads or Goroutines.
This allows better use of multicore processors and faster execution
for tasks that are computationally intensive.

GOROUTINES
  - Lightweight threads managed by the Go runtime (not the OS).
  - Launched with the `go` keyword before a function call.
  - Thousands of Goroutines can run concurrently with minimal overhead.
  - Provides a strong concurrency model for writing concurrent programs quickly.

CHANNELS
  - Used for communication and synchronization between Goroutines.
  - Goroutines can send and receive data through channels safely.
  - Prevents race conditions by controlling how data is shared.

=================================================================================
CONCURRENCY CONSEQUENCES
=================================================================================

Concurrent programs introduce risks that need to be understood and handled.

Race Condition
  Two or more goroutines access the same variable simultaneously without proper
  synchronization — one writes while another reads, producing unpredictable,
  inconsistent results. Hard to reproduce and dangerous in production.

Deadlock
  Goroutines wait on each other indefinitely for a resource neither will release.
  Classic cause: goroutine A holds mutex 1 waiting for mutex 2, while goroutine B
  holds mutex 2 waiting for mutex 1. Go runtime auto-detects this and halts
  the program with a clear error.

Starvation
  A goroutine never gets access to the resource it needs because higher-priority
  goroutines keep taking it first. Critical tasks get delayed, response time spikes,
  and the app feels slow or unresponsive.

Livelock
  Similar to deadlock, but goroutines keep actively changing state in response to
  each other — never actually stopping, yet never making real progress.
  CPU gets consumed with no productive output.

Resource Contention
  Many goroutines compete for the same resource simultaneously, creating a
  performance bottleneck. Example: a full channel with no reader blocks the
  sending goroutine, reducing overall app throughput.

=================================================================================
SOLUTIONS
=================================================================================

Mutex Protection      → use sync.Mutex to ensure only one goroutine accesses
                        a shared resource at a time

Channel Communication → use channels for inter-goroutine communication instead
                        of sharing memory directly — safer and more idiomatic Go

WaitGroup Sync        → use sync.WaitGroup to synchronize goroutine lifecycles,
                        ensuring main() waits for all goroutines to finish

Consistent Locking    → always acquire mutexes in the same order across all
                        goroutines to prevent deadlock by design

Testing               → run `go run -race` to detect race conditions during development

=================================================================================
*/

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

// calling variadic function by unpacking a slice with ellipsis (...)
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

// declaring function for interface shapes
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

// declaring method to circle
func (c circle) volume() float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(c.radius, 3)
}

func factorial(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	fmt.Printf("Factorial of %d is %d\n", n, result)
}

func firstProcess(index int) {
	fmt.Println("First process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("First process func ended")
}

func secondProcess(index int) {
	fmt.Println("Second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("Second process func ended")
}

func firstProcess1(index int) {
	fmt.Println("First Process func starter")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("First process func ended")
}

func secondProcess1(index int) {
	fmt.Println("Second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("Second process func ended")
}

// declaring goroutine for email nontification
func sendEmailAsync(userID int, message string) {
	// Stimulate sending an email (in real-life scenario)
	time.Sleep(2 * time.Second)
	fmt.Printf("Email notification sent to user %d: %s\n", userID, message)
}

// declaring gorutine for image processing
func processImage(imageURL string) {
	fmt.Printf("Processing image: %s\n", imageURL)
	// Simulate image processing (replace this with actual image processing code)
	time.Sleep(3 * time.Second)
	fmt.Printf("Image processing completed: %s\n", imageURL)
}

func scheduleTask(task func()) {
	go task()
}

// declaring function for goroutine file download
func downloadFile(url string, destination string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error downloading file from %s: %s\n", url, err)
		return
	}
	defer response.Body.Close()

	file, err := os.Create(destination)
	if err != nil {
		fmt.Printf("Error creating file: %s: %s\n", destination, err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Printf("Error writing to file %s: %s\n", destination, err)
		return
	}

	fmt.Printf("Downloaded file from %s to %s\n", url, destination)
}

// declaring channels
func introduce(student string, c chan string) {
	result := fmt.Sprintf("Hai, my name is %s", student)

	c <- result
}

// declaring channels for anonymous function
func printChannel(c chan string) {
	fmt.Println(<-c)
}

func printChannel1(c3 <-chan string) {
	fmt.Println(<-c3)
}

func introduce1(student string, c chan<- string) {
	result := fmt.Sprintf("Hai, my name is %s", student)
	c <- result
}

func calculateSumOfSquares(numbers []int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, num := range numbers {
		sum += num * num
	}
	ch <- sum
}

func countWords(filename string, ch chan<- int) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", filename, err)
		ch <- 0
		return
	}

	words := strings.Fields(string(content))
	ch <- len(words)
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("Defer func starts to execute")
}

func processFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lines, nil
}

func performTask(taskName string) {
	fmt.Printf("Task %s started. \n", taskName)
	defer fmt.Printf("Task: %s completed. \n", taskName)
}

func validPassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("password has to have more than 4 characters")
	}
	return "Valid password", nil
}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func validPassword1(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("Password has to have more than 4 characters")
	}
	return "valid password", nil
}

func divide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Division by zero occured.")
		}
	}()

	if b == 0 {
		panic("Division by zero")
	}

	return a / b, nil
}
