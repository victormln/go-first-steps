package main

import (
	"fmt"
	"strings"

	"./name"
	"./numbers"
	"./structure"
)

const welcomeMessage string = "Welcome %s!"

func main() {
	firstName := name.GetName()
	fmt.Printf(welcomeMessage, firstName)

	age := numbers.GetAge()
	fmt.Println(age)

	myNameInJapanese := name.GenerateNameInJapanese()
	fmt.Println("Victor in japanese: ", myNameInJapanese)

	var resultOfSum = numbers.Sum(1, 9)
	fmt.Println(resultOfSum)

	int32, int16, int64 := numbers.GetVariables()
	fmt.Println(int16, int32, int64)

	float32, float64 := numbers.GetFloat()
	fmt.Println(float32, float64)

	var text = "testing"
	fmt.Println("Extract letter from", text, "(", string("testing"[1]), ")")
	fmt.Println(len(string("testing")))

	getArray()

	getSlice()

	structure.IfTest()

	structure.ForTest()

	strings2()

	structure.SwitchTest()
}

func getArray() {
	var typeOfCar [2]string
	randomNumbers := [3]int{5, 3, 12}
	fmt.Println(randomNumbers)
	typeOfCar[0] = "electric"
	typeOfCar[1] = "gasoil"
	fmt.Println(typeOfCar)
}

func getSlice() {
	var slice []string
	slice = append(slice, "hola", "que", "tal")
	fmt.Println(slice)
}

func strings2() {
	message := "Hoooola, que tal, no paras de decir y tal y tal?"
	fmt.Println(strings.ToUpper(message))
	fmt.Println(strings.ToLower(message))
	fmt.Println(strings.Replace(message, "tal", "ase", 0))
	fmt.Println(strings.Replace(message, "tal", "ase", 1))
	fmt.Println(strings.Replace(message, "tal", "ase", -1))
	messageInArray := strings.Split(message, ",")
	fmt.Println(messageInArray)
	fmt.Println("Array length: ", len(messageInArray))
}
