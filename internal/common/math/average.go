package math

import "sort"

func Mean(data []float64) float64 {

	var sum float64
	for _, item := range data {
		sum += item
	}

	return sum / float64(len(data))
}

func Min(data []float64) float64 {

	sort.Float64s(data)

	return data[0]
}

func Max(data []float64) float64 {

	sort.Float64s(data)

	return data[len(data)-1]
}
