package vector

// TODO: Should return an error in the case where the slice is empty
func Centroid(vs ...T) T {
	count := len(vs)
	xSum := 0.0
	ySum := 0.0
	for _, v := range vs {
		xSum += v.X()
		ySum += v.Y()
	}
	return New(xSum/float64(count), ySum/float64(count))
}
