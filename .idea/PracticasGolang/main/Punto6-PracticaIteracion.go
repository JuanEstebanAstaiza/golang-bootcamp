package main

import (
	"fmt"
	"strings"
)

func fizzBuzz(n int) string {
	switch {
	case n%3 == 0 && n%5 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprint(n)
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		resultado := fizzBuzz(i)
		fmt.Println(strings.Repeat("-", 20))
		fmt.Printf("Número: %d\nResultado: %s\n", i, resultado)
	}
}
