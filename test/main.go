package main

import "fmt"

// type person struct {
// 	name string
// 	age  int
// }

// func main() {
// 	// 定义一个 person 变量 a
// 	a := person{name: "Tom", age: 20}

// 	// 赋值操作 b = a
// 	b := a
// 	b.age = 30
// 	fmt.Println(a, b, &a, &b)
// 	fmt.Printf("a: %v\n", &a)

// 	// // 赋值操作 b = &a
// 	// c := &a
// 	// c.age = 40
// 	// fmt.Printf("a: %v, c: %v\n", a, *c)
// }

func main() {
	a := []int{1, 2, 3}
	b := a
	b[0] = 4
	fmt.Println(a, b)
	fmt.Printf("a: %v, b: %v\n", &a, &b)
}
