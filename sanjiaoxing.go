package main

import (
	"fmt"
)

func main() {
	for i := 10; i >= 0; i-- {
		for l := 0; l < 10-i; l++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println("20210504")
	}
}
