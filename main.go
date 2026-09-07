package main

import "fmt"

func main() {
	Loop()
}
func Loop() {
	for i := range 99 {
		fmt.Println("Number:", i)
	}
}
