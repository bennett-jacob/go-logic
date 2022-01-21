package main

import (
	"fmt"

	"github.com/bennett-jacob/go-logic"
)

func main() {
	fmt.Printf("false, false, false = %t\n", logic.None(false, false, false))
	fmt.Printf("true, true, false = %t\n", logic.None(true, true, false))
	fmt.Printf("true, true, true = %t\n", logic.None(true, true, true))
}
