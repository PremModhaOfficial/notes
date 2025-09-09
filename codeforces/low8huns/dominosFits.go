package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var w string
	// Read the weight from standard input
	_, err := fmt.Scanln(&w)
	if err != nil {
		return
	}

	nums := strings.Split(w, " ")
	area, err := strconv.ParseInt(nums[0], 1, 10)
	side, err := strconv.ParseInt(nums[1], 1, 10)
	area *= side

	// root of 2 fastes
	sqr := int64(1)
	for sqr <= area {
		sqr *= 2
	}

	fmt.Println(sqr / 2)
}
