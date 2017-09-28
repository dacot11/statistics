package normalization

import "github.com/dacot11/statistics/sliceutil"

func Minmax(value float64, set []float64, newMin float64, newMax float64) float64 {

	var min float64 = sliceutil.Min(set)
	var max float64 = sliceutil.Max(set)
	var newValue float64

	newValue = (((value - min) / (max - min)) * (newMax - newMin)) + newMin

	return newValue
}

func Zscore(value float64, set []float64, standardDeviation float64) float64 {

	return (value - sliceutil.Mean(set)) / standardDeviation
}
