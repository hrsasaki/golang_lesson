package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type PackageInfo struct {
	Deps []string
}

func main() {
	var result PackageInfo
	json.NewDecoder(os.Stdin).Decode(&result)
	fmt.Fprint(os.Stdout, result.Deps)
}
