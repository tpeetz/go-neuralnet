package probabilistic

import "testing"

func Test_normalize(t *testing.T) {
	var testArray1 = []float64{3.0, 4.0}
	type args struct {
		input *[]float64
	}
	tests := []struct {
		name string
		args args
	}{
		{"using testArray1", args{&testArray1}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normalize(tt.args.input)
		})
	}
}
