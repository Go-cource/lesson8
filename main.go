package main

import "fmt"

func main() {
	planets := [8]string{
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран",
		"Нептун",
	}
	fmt.Println(planets)
	// for i := 0; i < len(planets); i++ {
	// 	fmt.Println(planets[i])
	// }
	for index, value := range planets {
		fmt.Printf("Элемент с индексом: %d - %s\n", index, value)
	}
}
