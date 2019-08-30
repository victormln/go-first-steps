package numbers

import "fmt"

const insertAgeMessage = "\nHow old are you?: "

// GetAge returns the user age
func GetAge() int {
	var age int
	fmt.Print(insertAgeMessage)
	fmt.Scanf("%d", &age)

	return age
}

// GetVariables return 3 ints
func GetVariables() (int, int16, int64) {
	return 212, 1, 12312412412412
}

// GetFloat returns two floats
func GetFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

// Sum return the result of the sum of two numbers
func Sum(number1 int, number2 int) int {
	return number1 + number2
}
