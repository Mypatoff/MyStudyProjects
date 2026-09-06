package main

import "fmt"

func main() {
	anLoop()
}
func anLoop() {
	for i := range 99 {
		fmt.Println("Number:", i)
	}
}
