package commands

import (
	"fmt"
)

// Normalize command triggers a normalization process.
func Normalize() {
	fmt.Println("1) Min-max")
	fmt.Println("2) Z-socre")

	var option string
	fmt.Scanln(&option)

	fmt.Println("TODO")
}
