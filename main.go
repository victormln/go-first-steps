package main

import (
	"fmt"
	"strings"

	/*"./array"
	"./name"
	"./numbers"
	"./structure"*/
	"./maps"
)

const welcomeMessage string = "Welcome %s!"

func main() {
	fmt.Println(maps.GetMap())
	var name = "victor"

	fmt.Println(maps.GetAgeFromName(name))
	/*
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

		array.GetArray()

		array.GetSlice()

		structure.IfTest()

		structure.ForTest()

		strings2()

		structure.SwitchTest()
	*/
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
