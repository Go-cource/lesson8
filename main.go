package main

import "fmt"

func main() {
	a := 10
	var p *int //Указатель на int
	p = &a
	fmt.Println(a, p)
	*p = 15
	fmt.Println(a)
	//Нельзя делать указатель на const
	// const c = 5
	// p1 := &"asdasd"
	// p2 := &c

}
