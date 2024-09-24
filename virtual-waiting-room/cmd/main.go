package main

import (
	"fmt"

	"github.com/gabrielmaurici/eventim-simulation/pkg/token"
)

func main() {
	tokme, err := token.GenerateAccessToken()
	if err != nil {
		panic("erro")
	}
	fmt.Println("token: " + tokme)
}
