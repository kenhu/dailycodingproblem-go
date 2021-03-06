package day446

import "testing"

// nolint
var testcases = []struct {
	n           int32
	isPowerFour bool
}{
	{16, true},
	{15, false},
	{0, false},
}

func TestIsPowerFourBrute(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := IsPowerFourBrute(tc.n); result != tc.isPowerFour {
			t.Errorf("For n=%d, expected %v, got %v", tc.n, tc.isPowerFour, result)
		}
	}
}

func BenchmarkIsPowerFourBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IsPowerFourBrute(tc.n)
		}
	}
}

func TestIsPowerFourFaster(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := IsPowerFourFaster(tc.n); result != tc.isPowerFour {
			t.Errorf("For n=%d, expected %v, got %v", tc.n, tc.isPowerFour, result)
		}
	}
}

func BenchmarkIsPowerFourFaster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IsPowerFourFaster(tc.n)
		}
	}
}
