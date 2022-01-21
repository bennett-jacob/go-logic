package main

import (
	"fmt"

	"github.com/bennett-jacob/go-logic"
)

func main() {
	fmt.Printf("false, false = %t\n", logic.Xor(false, false))
	fmt.Printf("true, false = %t\n", logic.Xor(true, false))
	fmt.Printf("false, true = %t\n", logic.Xor(false, true))
	fmt.Printf("true, true = %t\n", logic.Xor(true, true))
}
