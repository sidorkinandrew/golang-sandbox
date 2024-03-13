package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)

	fmt.Println("It is", a/30, "hours", 2*(a%30), "minutes.")
}
