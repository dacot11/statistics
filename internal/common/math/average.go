package math

import "sort"

// Mean calculates the mean in an array of numbers.
func Mean(data []float64) float64 {
	var sum float64
	for _, item := range data {
		sum += item
	}

	return sum / float64(len(data))
}

// Min finds the minimum in an array of numbers.
func Min(data []float64) float64 {

	sort.Float64s(data)

	return data[0]
}

// Max finds the maximum in an array of numbers.
func Max(data []float64) float64 {

	sort.Float64s(data)

	return data[len(data)-1]
}
