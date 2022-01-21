package main

import (
	"fmt"

	"github.com/bennett-jacob/go-logic"
)

func main() {
	fmt.Printf("false, false, false = %t\n", logic.Any(false, false, false))
	fmt.Printf("true, false, false = %t\n", logic.Any(true, false, false))
	fmt.Printf("true, true, true = %t\n", logic.Any(true, true, true))
}
