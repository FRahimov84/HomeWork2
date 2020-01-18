package main

import "testing"

func Test_sumOfBonuses(t *testing.T) {

	tests := []struct {
		name  string
		sales []int64
		want  int64
	}{
		// TODO: Add test cases.
		{"test: have not a bonus",[]int64{9_000, 6_000, 10_000, 10_000},0},
		{"test: have a bonus",[]int64{19_000, 16_000, 10_000, 10_000},750},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := sumOfBonuses(tt.sales); gotSum != tt.want {
				t.Errorf("%s: got: %v, want: %v",tt.name, gotSum, tt.want)
			}
		})
	}
}