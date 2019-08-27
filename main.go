package main

import "fmt"

const insertNameMessage = "Insert your name:"
const insertAgeMessage = "How old are you?:"
const welcomeMessage string = "Welcome %s!"

func main() {
	name := getName()
	fmt.Printf(welcomeMessage, name)

	age := getAge()
	fmt.Scanf("%d", &age)
	fmt.Println(age)
}

func getName() string {
	var name string
	name = "Default name"

	fmt.Print(insertNameMessage)
	fmt.Scanf("%s", &name)

	return name
}

func getAge() int {
	var age int = 0
	fmt.Print(insertAgeMessage)
	fmt.Scanf("%d", &age)

	return age
}
