package main

import (
	"fmt"

	"github.com/dacot11/statistics/cmd/statistics/commands"
)

func main() {
	option := requestStatisticsOption()

	switch option {
	case "1":
		commands.CalculateDistance()
	case "2":
		commands.Normalize()
	default:
		fmt.Println("Invalid option")
	}
}

func requestStatisticsOption() string {

	fmt.Println("Please select an option:")
	fmt.Println("1) Distance")
	fmt.Println("2) Normalization")

	var option string
	fmt.Scanln(&option)

	return option
}
