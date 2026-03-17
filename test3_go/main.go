package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
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

	//calling variadic functions
	studentLists := print("Airell", "Nanda", "Mailo", "Schannel", "Marco")

	fmt.Printf("%v", studentLists)
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
