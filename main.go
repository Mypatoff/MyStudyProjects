package main

import (
	"fmt"
	"time"
)

func main() {
	// basicIf(18)
	for i := range 5 {
		goRoutine(i)
		time.Sleep(time.Second)
	}
	for i := 5; i <= 10; i++ {
		goRoutine(i)
		time.Sleep(time.Second)
	}
}

//	func basicIf(n int) {
//		if n%2 == 1 {
//			fmt.Println("Odd")
//		} else {
//			fmt.Println("Not odd")
//		}
//	}
func goRoutine(i int) {
	fmt.Println(i)
}
