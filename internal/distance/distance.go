package distance

import "math"

func Minkowski(h float64, n int, firstObject []float64, secondObject []float64) float64 {

	var sum float64

	for i := 0; i < n; i++ {
		sum += math.Pow(math.Abs(float64(firstObject[i]-secondObject[i])), h)
	}

	return math.Pow(sum, 1/h)
}

func Supremum(n int, firstObject []float64, secondObject []float64) float64 {

	var maxDistance float64 = 0

	for i := 0; i < n; i++ {
		var distance float64 = math.Abs(float64(firstObject[i] - secondObject[i]))

		if distance > maxDistance {
			maxDistance = distance
		}
	}

	return maxDistance
}

func Cosine(n int, firstObject []float64, secondObject []float64) float64 {

	var product float64 = cartesianProduct(n, firstObject, secondObject)

	var firstEuclideanNorm float64 = euclideanNorm(n, firstObject)

	var secondEuclideanNorm float64 = euclideanNorm(n, secondObject)

	return product / (firstEuclideanNorm * secondEuclideanNorm)

}

func cartesianProduct(n int, firstObject []float64, secondObject []float64) float64 {

	var product float64

	for i := 0; i < n; i++ {
		product += float64(firstObject[i] * secondObject[i])
	}

	return product
}

func euclideanNorm(n int, object []float64) float64 {

	var euclideanNorm float64 = 0

	for i := 0; i < n; i++ {
		euclideanNorm += math.Pow(object[i], 2)
	}

	return math.Sqrt(euclideanNorm)
}
