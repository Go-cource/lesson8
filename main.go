package main

import "fmt"

func terraform(planets [9]string) {
	//ЗДЕСЬ СВОЙ PLANETS [8]string
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
}

func main() {
	planets := [...]string{
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран",
		"Нептун",
	}

	terraform(planets) //[new Меркурий, new etc]
	fmt.Println(planets)
}
