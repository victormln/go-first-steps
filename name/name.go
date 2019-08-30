package name

import "fmt"

const insertNameMessage = "Insert your name: "

// GetName obtiene y retorna el nombre
func GetName() string {
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
