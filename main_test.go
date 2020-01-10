package main

import "testing"

func Test_sumOfBonuses(t *testing.T) {

	tests := []struct {
		name  string
		sales []int
		want  float64
	}{
		// TODO: Add test cases.
		{"test without oversum",[]int{9_000, 6_000, 5_000, 1_000},0},
		{"test with one oversum",[]int{9_000, 16_000, 5_000, 1_000},800},
		{"test with oversums",[]int{90_000, 16_000, 50_000, 10_000},8_300},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := sumOfBonuses(tt.sales); gotSum != tt.want {
				t.Errorf("sumOfBonuses() = %v, want %v", gotSum, tt.want)
			}
		})
	}
}