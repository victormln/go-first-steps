package structure

import "fmt"

// IfTest test if a number is an even number or not
func IfTest() {
	var number = 0
	fmt.Println("Ingresa un numero del 1 al 10")
	fmt.Scanf("%d", &number)
	if number%2 == 0 {
		fmt.Println("Numero par")
	} else {
		fmt.Println("Numero impar")
	}

	// Set again number2 variable to 3 and compare it to enter inside
	if number2 := 3; number2 == 3 {
		fmt.Println("Dentro!")
	}
}

// ForTest test three ways to do a loop
func ForTest() {
	StandardFor()

	LoopWithoutInitializeIntoIt()

	InfiniteLoopWithBreak()
}

// StandardFor do a basic for
func StandardFor() {
	for counter := 0; counter < 3; counter++ {
		fmt.Print(counter, ", ")
	}
	fmt.Println()
}

// LoopWithoutInitializeIntoIt loop without inicialize a variable
func LoopWithoutInitializeIntoIt() {
	c := 6
	for ; c > 0; c = c - 2 {
		fmt.Print(c, " ")
	}
}

// InfiniteLoopWithBreak break an infinite for
func InfiniteLoopWithBreak() {
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

// SwitchTest if the number is 1 or not
func SwitchTest() {
	var number = 0
	fmt.Print("Insert a number: ")
	fmt.Scanf("%d", &number)

	switch number {
	case 1:
		fmt.Println("El numero es uno")
	default:
		fmt.Println("Numero no es 1")
	}

	switch {
	case number%2 == 0:
		fmt.Println("El numero es par")
	default:
		fmt.Println("El numero es impar")
	}
}
