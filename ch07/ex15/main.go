package main

import (
	"fmt"
	"os"
)

func main() {
	expr, err := Parse(os.Args[1], true)
	if err != nil {
		panic(fmt.Sprintf("invalid expr: %s", os.Args[1]))
	}
	fmt.Printf("result: %g\n", expr.Eval(parseEnv))
}
