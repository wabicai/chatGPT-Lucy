package main

import "fmt"

type myStruct struct {
	ee int
}

func main() {
	var a myStruct = myStruct{}
	var b *myStruct = &a
	a.ee = 2
	fmt.Println(a.ee, (*b).ee, &a, a, &b)
	// fmt.Println(*a)
}
