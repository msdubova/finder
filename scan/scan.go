package scan

import (
	"bufio"
	"fmt"
	"os"
)

func Scan() []string {
	var file, error = os.Open("File1.txt")
	if error != nil {
		fmt.Println("Помилка відкриття файлу")
	}

	defer file.Close()

	var text = bufio.NewScanner(file)
	var slice []string
	for text.Scan() {
		var line = text.Text()
		slice = append(slice, line)
	}
	fmt.Println("Довжина зрізу:", len(slice))
	return slice

}
