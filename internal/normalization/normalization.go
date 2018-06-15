package normalization

import "github.com/dacot11/statistics/internal/common/math"

// Minmax normalizes a set using the Minmax strategy.
func Minmax(value float64, set []float64, newMin float64, newMax float64) float64 {

	min := math.Min(set)
	max := math.Max(set)
	var newValue float64

	newValue = (((value - min) / (max - min)) * (newMax - newMin)) + newMin

	return newValue
}

// Minmax normalizes a set using the Zscore strategy.
func Zscore(value float64, set []float64, standardDeviation float64) float64 {

	return (value - math.Mean(set)) / standardDeviation
}
