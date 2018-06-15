package distance

import "math"

// Minkowski calculates the Minkowski distance between to objects.
func Minkowski(h float64, n int, firstObject []float64, secondObject []float64) float64 {

	var sum float64

	for i := 0; i < n; i++ {
		sum += math.Pow(math.Abs(float64(firstObject[i]-secondObject[i])), h)
	}

	return math.Pow(sum, 1/h)
}

// Supremum calculates the Supremum distance between to objects.
func Supremum(n int, firstObject []float64, secondObject []float64) float64 {

	var maxDistance float64

	for i := 0; i < n; i++ {
		distance := math.Abs(float64(firstObject[i] - secondObject[i]))

		if distance > maxDistance {
			maxDistance = distance
		}
	}

	return maxDistance
}
