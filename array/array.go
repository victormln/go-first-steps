package array

import "fmt"

// GetArray create and array and set values
func GetArray() {
	var typeOfCar [2]string
	randomNumbers := [3]int{5, 3, 12}
	fmt.Println(randomNumbers)
	typeOfCar[0] = "electric"
	typeOfCar[1] = "gasoil"
	fmt.Println(typeOfCar)
}

// GetSlice append to array
func GetSlice() {
	var slice []string
	slice = append(slice, "hola", "que", "tal")
	fmt.Println(slice)
}
