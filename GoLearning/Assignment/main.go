package main

import (
	"fmt"
)

func main() {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range list {
		if list[i]%2 == 0 {
			fmt.Println(list[i], "is even")
		} else {
			fmt.Println(list[i], "is odd")
		}
	}
}
