package main

import (
	"fmt"
)

type student struct {
	name  string
	age   int
	class string
}

func (s *student) GetName() string {
	return s.name
}

type batch struct {
	batchproperties student
	batchname       string
}

type score struct {
	student student
	score   int
}

func (s *score) LulusAtauGa() string {
	if s.score < 10 {
		return "tidak lulus"
	} else {
		return "lulus"
	}
}

func main() {
	student1 := student{
		name:  "student1",
		age:   20,
		class: "class1",
	}
	fmt.Println(student1.GetName())

	score1 := score{student: student1, score: 1}
	fmt.Println(score1.LulusAtauGa())

	//ngcVar1()
	//fmt.Println()
	//ngcVar2()
	//fmt.Println()
	//ngcCliProgram()
	//ngcArrayAndSlice()

	//text := "please welcome"
	//names := []string{"Gopher", "Gopher2", "Gopher3", "Gopher4", "Gopher5"}
	//myName := CallMyname("Gopher")
	//fmt.Println(myName)
	//
	//theirName := CallTheirNames(text, "Gopher", "Gopher2", "Gopher3", "Gopher4", "Gopher5")
	//fmt.Println(theirName)
	//
	//names1 := CallNames(text, names)
	//fmt.Println(names1)
	//texted := fmt.Sprintf("%s %s", text, names[1])
	//fmt.Println(texted)

	//area, keliling, err := HitungAreaDanKelilingSegi4(-1, 5)
	//if err != nil {
	//	return
	//}
	//fmt.Println("area: ", area)
	//fmt.Println("keliling: ", keliling)
	//fmt.Println("error: ", err)

	//data := CallMyname("jamie")
	//fmt.Println(data)

	//var a, b, c int = 1, 2, 0
	//
	//perkalian(a, b, &c)
	//fmt.Println(c)

	//var a, b, respCode, respMessage string = "a", "", "", ""
	//fmt.Println(respCode)
	//fmt.Println(respMessage)
	//RespHandler(a, b, &respCode, &respMessage)
	//fmt.Println(respCode, respMessage)

}

func RespHandler(a, b string, respCode, respMessage *string) {
	if len(a) < 1 {
		*respCode = "error"
		*respMessage = "Data a kurang"
		return
	}

	if len(b) < 1 {
		*respCode = "error"
		*respMessage = "Data b kurang"
		return
	}
}

func HitungAreaDanKelilingSegi4(panjang, lebar int) (area, keliling int, err error) {
	if panjang < 1 {
		err = fmt.Errorf("panjangnya salah")
		return
	}

	if lebar < 1 {
		err = fmt.Errorf("lebarnya salah")
		return
	}

	area = panjang * lebar
	keliling = 2 * (panjang + lebar)
	return
}

func CallMyname(name string) (output string) {
	if len(name) < 1 {
		return ""
	}

	output = "Haloo " + name
	return
}

func perkalian(a, b int, c *int) {
	*c = a * b
	return
}

func CallNames(text string, names []string) (output string) {
	output = text
	for _, name := range names {
		output = output + " " + name
	}
	return
}

func CallTheirNames(text string, names ...string) (output string) {
	output = text
	for _, name := range names {
		output = output + " " + name
	}
	return
}

func ngcVar1() {
	var myNum int32 = 50
	fmt.Println(myNum)

	var myNum2 float32 = 51.0
	fmt.Println(myNum2)

	var myNumStr string = "50"
	fmt.Println(myNumStr)
}

func ngcVar2() {
	x := 5
	y := 10
	z := x + y
	fmt.Println(z)
}

func ngcCliProgram() {
	var name string
	fmt.Println("masukkan nama: ")
	fmt.Scanln(&name)

	fmt.Println("hello " + name)
}

func ngcArrayAndSlice() {
	people := []string{"Walt", "Jesse", "Skyler", "Saul"}
	fmt.Println(len(people))

	people = append(people, "Hank", "Marie")
	fmt.Println(len(people))
	fmt.Println(people)
}
