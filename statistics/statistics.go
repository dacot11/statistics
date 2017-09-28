package main

import (
	"fmt"
	"github.com/dacot11/statistics/distance"
	"strconv"
	"strings"
)

func main() {

	var option string = ""

	option = requestStatisticsOption(option)

	if option == "1" {
		option = requestDistanceOption(option)
	} else if option == "2" {
		option = requestNormalizationOption(option)
	} else {
		fmt.Println("Invalid option")
	}
}

func requestStatisticsOption(option string) string {

	fmt.Println("Please select an option:")
	fmt.Println("1) Distance")
	fmt.Println("2) Normalization")

	fmt.Scanln(&option)

	return option
}

func requestDistanceOption(option string) string {

	fmt.Println("1) Minkowski")
	fmt.Println("2) Supremum")

	fmt.Scanln(&option)

	if option == "1" {
		processMinkowski()
	} else if option == "2" {
		processSupremum()
	} else {
		fmt.Println("Invalid option")
	}

	return option
}

func processMinkowski() {

	fmt.Println("Enter H:")
	rawH := ""
	fmt.Scanln(&rawH)
	h, _ := strconv.ParseFloat(rawH, 64)

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

func requestNormalizationOption(option string) string {

	fmt.Println("1) Min-max")
	fmt.Println("2) Z-socre")

	fmt.Scanln(&option)

	return option
}
