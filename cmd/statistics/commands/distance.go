package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dacot11/statistics/internal/distance"
)

// CalculateDistance triggers a distance calculation process.
func CalculateDistance() {
	fmt.Println("1) Minkowski")
	fmt.Println("2) Supremum")

	var option string
	fmt.Scanln(&option)

	switch option {
	case "1":
		processMinkowski()
	case "2":
		processSupremum()
	default:
		fmt.Println("Invalid option")
	}
}

func processMinkowski() {
	fmt.Println("Enter H:")
	var rawH string
	fmt.Scanln(&rawH)
	h, _ := strconv.ParseFloat(rawH, 64)

	fmt.Println("Enter N:")
	var rawN string
	fmt.Scanln(&rawN)
	n, _ := strconv.ParseInt(rawN, 0, 4)

	fmt.Println("Enter first set (comma separated values):")
	var setData string
	fmt.Scanln(&setData)
	set1 := buildFloatSlice(setData)

	fmt.Println("Enter first set (comma separated values):")
	fmt.Scanln(&setData)
	set2 := buildFloatSlice(setData)

	fmt.Printf("Minkoski Distance with h = %f between %v and %v is: %f \n", h, set1, set2, distance.Minkowski(h, int(n), set1, set2))

}

func processSupremum() {
	fmt.Println("Enter N:")
	rawN := ""
	fmt.Scanln(&rawN)
	n, _ := strconv.ParseInt(rawN, 0, 4)

	fmt.Println("Enter first set (comma separated values):")
	setData := ""
	fmt.Scanln(&setData)
	set1 := buildFloatSlice(setData)

	fmt.Println("Enter first set (comma separated values):")
	fmt.Scanln(&setData)
	set2 := buildFloatSlice(setData)

	fmt.Printf("Supremum Distance between %v and %v is: %f \n", set1, set2, distance.Supremum(int(n), set1, set2))

}

func buildFloatSlice(rawData string) []float64 {
	floatSlice := []float64{}

	stringSlice := strings.Split(rawData, ",")

	for _, item := range stringSlice {
		floatItem, _ := strconv.ParseFloat(item, 64)
		floatSlice = append(floatSlice, floatItem)
	}

	return floatSlice
}
