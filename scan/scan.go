package scan

import "fmt"

func Scan() {
	var text string

	fmt.Println("Введіть текст для пошуку")
	fmt.Scanln(&text)

	fmt.Printf(text)
}
