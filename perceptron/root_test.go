package perceptron

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		input  int
		hidden []int
		output int
	}
	tests := []struct {
		name string
		args args
		want *Network
	}{
		{"", args{2, []int{2}, 2}, &Network{}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.input, tt.args.hidden, tt.args.output); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNetwork_Propagate(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name    string
		network *Network
		args    args
		want    []float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			network := &Network{}
			if got := network.Propagate(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Network.Propagate() = %v, want %v", got, tt.want)
			}
		})
	}
}
