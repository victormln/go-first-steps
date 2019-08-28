package main

import "fmt"

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

func sum(number1 int, number2 int) int {
	return number1 + number2
}
