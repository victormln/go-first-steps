package name

import "fmt"

const insertNameMessage = "Insert your name: "

// AskForName obtiene y retorna el nombre
func AskForName() string {
	var name string
	name = "Default name"

	fmt.Print(insertNameMessage)
	fmt.Scanf("%s", &name)

	return name
}

// GenerateNameInJapanese the name "Victor" in Japanese to test UTF8
func GenerateNameInJapanese() string {
	return "ビクター"
}
