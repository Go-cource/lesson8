package main

import "fmt"

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
	terr := planets[4:]
	fmt.Println(terr)

}
