package main

import (
	"fmt"
)

func main() {
	
	var a int = 42
	var b *int = &a
	fmt.Println("b is a pointer", a, b)	
	fmt.Println("use * for getting value form pointer", a, *b)

	fmt.Println("the pointer addresses of both vars", &a, b)	

	*b = 14
	fmt.Println("change a value by setting b that is the same pointer", a, b)

	
}