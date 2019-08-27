package main

import "fmt"

const insertNameMessage = "Insert your name:"
const insertAgeMessage = "How old are you?:"
const welcomeMessage string = "Welcome %s!"

func main() {
	var name string
	name = "Default name"
	var age = 0
	withoutType := "Variable is set but without setting the type :)"
	fmt.Print(insertNameMessage)
	fmt.Scanf("%s", &name)
	fmt.Printf(welcomeMessage, name)
	fmt.Print(insertAgeMessage)
	fmt.Scanf("%d", &age)
	fmt.Println(age, withoutType)
}
