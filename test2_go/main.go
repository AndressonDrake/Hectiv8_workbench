package main

import (
	"fmt"
)

func main() {

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

}

func greet(name, address string) {
	fmt.Println("Hello there! My name is", name)
	fmt.Println("I live in", address)
}
