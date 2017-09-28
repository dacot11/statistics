package main

import (
	"fmt"
	"github.com/dacot11/statistics/distance"
	"github.com/dacot11/statistics/normalization"
)

func main() {

	fmt.Printf("Minkoski Distance with h = 1 is: %f \n", distance.MinkowskiDistance(1, 4, []int{22, 1, 42, 10}, []int{20, 0, 36, 8}))

	fmt.Printf("Minkoski Distance with h = 2 is: %f \n", distance.MinkowskiDistance(2, 4, []int{22, 1, 42, 10}, []int{20, 0, 36, 8}))

	fmt.Printf("Minkoski Distance with h = 3 is: %f \n", distance.MinkowskiDistance(3, 4, []int{22, 1, 42, 10}, []int{20, 0, 36, 8}))

	fmt.Printf("Supremum Distance is: %f \n", distance.SupremumDistance(4, []int{22, 1, 42, 10}, []int{20, 0, 36, 8}))

	// Cosine distance
	var x []float64 = []float64{1.4, 1.6}
	var y []float64 = []float64{1.5, 1.7}
	fmt.Printf("Cosine Distance between %v and %v is: %f \n", x, y, distance.CosineDistance(2, x, y))

	var y1 []float64 = []float64{2, 1.9}
	fmt.Printf("Cosine Distance between %v and %v is: %f \n", x, y1, distance.CosineDistance(2, x, y1))

	var y2 []float64 = []float64{1.6, 1.8}
	fmt.Printf("Cosine Distance between %v and %v is: %f \n", x, y2, distance.CosineDistance(2, x, y2))

	var y3 []float64 = []float64{1.2, 1.5}
	fmt.Printf("Cosine Distance between %v and %v is: %f \n", x, y3, distance.CosineDistance(2, x, y3))

	var y4 []float64 = []float64{1.5, 1}
	fmt.Printf("Cosine Distance between %v and %v is: %f \n", x, y4, distance.CosineDistance(2, x, y4))

	// Normalization
	var set []float64 = []float64{13, 15, 16, 16, 19, 20, 20, 21, 22, 22, 25, 25, 25, 25, 30, 33, 33, 35, 35, 35, 35, 36, 40, 45, 46, 52, 70}
	var value float64 = 35

	// Min-max
	var newMin float64 = 0.0
	var newMax float64 = 1.0
	fmt.Printf("Min-max nomalization for value %f in %v with new min %f and new max %f is: %f \n", value, set, newMin, newMax, normalization.Minmax(value, set, newMin, newMax))

	// Z-score
	var sd = 12.94
	fmt.Printf("Z-score nomalization for value %f in %v with SD %f is: %f \n", value, set, sd, normalization.Zscore(value, set, sd))
}
