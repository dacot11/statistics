package sliceutil

func Mean(data []float64) float64 {

	var sum float64
	for _, item := range data {
		sum += item
	}

	return sum / float64(len(data))
}
