package main

import (
	"fmt"
	"strings"
)

func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func main() {
	// planets := []string{" Венера ", "Земля", " Марс"}
	// hyperspace(planets)

	// planets2 := []string{" Нептун ", "  Плутон "}
	// hyperspace(planets2)

	// fmt.Println(strings.Join(planets, ""))
	// fmt.Println(strings.Join(planets2, ""))
	// fmt.Printf("%T", &planets)
}
