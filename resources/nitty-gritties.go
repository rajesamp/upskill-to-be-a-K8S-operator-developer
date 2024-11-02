package main

import (
	"fmt"
	"math"
)

func main() {
	valsq := math.Sqrt(1)
	res := fmt.Sprintf("number is %g", valsq)
	fmt.Printf(
	func() string {
		if valsq == 1.0 {
			return "number is 1 problem\n"
			} 
			return res
	}(),
	
)
}

// used inline functions as an alternate to if else block
//key is to use a anonym function --> funct() and call it as () in line 17
// instead of if else - use return statment
// mistakes: using an else keyword and not using the return keyword and missing anony function ☎️ ()
// wrapping anony function call within fmt.Printf()
