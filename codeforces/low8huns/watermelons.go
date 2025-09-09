package main

import (
	"fmt"
)

func main() {
	var w int
	// Read the weight from standard input
	_, err := fmt.Scan(&w)
	if err != nil {
		return
	}

	// Condition: Weight must be even and greater than 2
	if w > 2 && w%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
