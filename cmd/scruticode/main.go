package main

import (
	"Scruticode/internal/core/functions"
	"Scruticode/internal/shared/constants"
	"fmt"
)

func main() {
	fmt.Print(constants.GetBanner())
	functions.ExecuteScan()
}
