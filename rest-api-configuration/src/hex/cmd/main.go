package main

import (
	"fmt"
	"restapi/v2/internal/adapters/core/arithmetic"
	"restapi/v2/internal/ports"
	"rest"
)

func main() {

	// ports 
	// var core ports.ArithmeticPort

	// core = arithmetic.NewAdapter()

	arithAdapter := arithmetic.NewAdapter()
	result, err := arithAdapter.Addition(1, 3)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

