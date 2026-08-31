package main

import "fmt"

func main() {
	basicIf(18)
}
func basicIf(n int) {
	if n%2 == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Not odd")
	}
}
