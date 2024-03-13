package main

import "fmt"

func main() {
	fmt.Print("Ivan", 27) // Ivan27
	fmt.Println()         // Ivan27

	fmt.Println("Ivan", 27) // Ivan 27

	fmt.Print(33, 27) // 33 27
}
