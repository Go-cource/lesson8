package main

import (
	"fmt"
	"strings"
)

func hyperspace(worlds []string) {
	//worlds []string - ссылка на тот же самый массив, что и []string{" Венера ", "Земля", " Марс"}
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func main() {
	planets := []string{" Венера ", "Земля", " Марс"} // где-то в память есть оригианльный массив, на который этот слайс ссылается
	hyperspace(planets)

	planets2 := []string{" Нептун ", "  Плутон "}
	hyperspace(planets2)

	fmt.Println(strings.Join(planets, ""))
	fmt.Println(strings.Join(planets2, ""))
	fmt.Printf("%p", &planets)
}
