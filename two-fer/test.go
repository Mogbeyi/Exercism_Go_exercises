package main

import (
	"fmt"
	"strings"
)

func main() {
	name := strings.Split("Emmnauel","")

	for i := range name {
		fmt.Println(name[i])
	}
}
