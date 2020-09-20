package probabilistic

import "math"

// PatternCell defines the pattern cell of the Probabilistic Neuronal Network.
type PatternCell struct {
	Weight  []float64
	Pattern interface{}
}

// PNN defines the Probabilistic Neuronal Network
type PNN struct {
	pattern []PatternCell
}

// AddPattern add pattern to probabilistic neuronal network.
func (pnn *PNN) AddPattern(vector []float64, pattern interface{}) {

}

// normalize the input vector
func normalize(input *[]float64) {
	var sum float64
	vector := *input
	for _, value := range *input {
		sum += value * value
	}
	factor := 1.0 / math.Sqrt(sum)
	for index := range vector {
		vector[index] = vector[index] * factor
	}
	input = &vector
}
