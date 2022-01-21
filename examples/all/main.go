package main

import (
	"fmt"

	"github.com/bennett-jacob/go-logic/logic"
)

func main() {
	fmt.Printf("false, false, false = %t\n", logic.All(false, false, false))
	fmt.Printf("true, true, false = %t\n", logic.All(true, true, false))
	fmt.Printf("true, true, true = %t\n", logic.All(true, true, true))
}
