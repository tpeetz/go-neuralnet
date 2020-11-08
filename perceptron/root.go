package perceptron

// Layer defines the hidden and outlayer of the perceptron network.
type Layer struct {
}

// Network defines the multilayer perceptron network.
type Network struct {
}

// New creates an instance of a multilayer perceptron network.
func New(input int, hidden []int, output int) *Network {
	network := &Network{}
	return network
}

// Propagate calculates the network output with the given input.
func (network *Network) Propagate(input []float64) []float64 {
	return nil
}
