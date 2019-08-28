package main

import (
	"fmt"
	"strings"
)

const insertNameMessage = "Insert your name: "
const insertAgeMessage = "\nHow old are you?: "
const welcomeMessage string = "Welcome %s!"

func main() {
	name := getName()
	fmt.Printf(welcomeMessage, name)

	age := getAge()
	fmt.Println(age)

	myNameInJapanese := generateNameInJapanese()
	fmt.Println("Victor in japanese: ", myNameInJapanese)

	var resultOfSum = sum(1, 9)
	fmt.Println(resultOfSum)

	int32, int16, int64 := getVariables()
	fmt.Println(int16, int32, int64)

	float32, float64 := getFloat()
	fmt.Println(float32, float64)

	var text = "testing"
	fmt.Println("Extract letter from", text, "(", string("testing"[1]), ")")
	fmt.Println(len(string("testing")))

	getArray()

	getSlice()

	ifTest()

	forTest()

	strings2()
}

func getName() string {
	var name string
	name = "Default name"

	fmt.Print(insertNameMessage)
	fmt.Scanf("%s", &name)

	return name
}

func getAge() int {
	var age int
	fmt.Print(insertAgeMessage)
	fmt.Scanf("%d", &age)

	return age
}

func generateNameInJapanese() string {
	return "ビクター"
}

func getVariables() (int, int16, int64) {
	return 212, 1, 12312412412412
}

func getFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
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

func sum(number1 int, number2 int) int {
	return number1 + number2
}

func ifTest() {
	var number = 0
	fmt.Println("Ingresa un numero del 1 al 10")
	fmt.Scanf("%d", &number)
	if number%2 == 0 {
		fmt.Println("Numero par")
	} else {
		fmt.Println("Numero impar")
	}

	if number2 := 3; number2 == 3 {
		fmt.Println("Dentro!")
	}
}

func forTest() {
	standardFor()

	loopWithoutInitializeIntoit()

	infiniteLoopWithBreak()
}

func standardFor() {
	for counter := 0; counter < 3; counter++ {
		fmt.Print(counter, ", ")
	}
	fmt.Println()
}

func loopWithoutInitializeIntoit() {
	c := 6
	for ; c > 0; c = c - 2 {
		fmt.Print(c, " ")
	}
}

func infiniteLoopWithBreak() {
	fmt.Println()
	s := 1000
	for {
		s -= 1000
		if s == 0 {
			fmt.Println("Ha acabado el for")
			break
		}
	}
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
