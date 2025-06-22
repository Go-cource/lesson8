package main

import "fmt"

func main() {
	planets := [...]string{ // Компилятор Go подсчитывает элементы
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран",
		"Нептун",
		"центавра",
	}
	fmt.Println(planets)

}
