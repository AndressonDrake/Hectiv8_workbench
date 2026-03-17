package main

import "fmt"

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

	//Maps: key-value stores implemented hash tables
	//Map is Go's built-in associative data type to stores key-value pairs with fast average-time lookups
	//Maps let you use more meaningful keys such as names, IDs, or other comparable values
}
