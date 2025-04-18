package main

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMinBasic(t *testing.T) {
	if IntMin(1, 2) != 1 {
		t.Error("IntMin(1, 2)!= 1")
	}
}

func TestIntMinTableDriven(t *testing.T) {
	test := []struct {
		a, b, result int
	}{
		{1, 2, 1},
		{2, 1, 1},
		{1, 1, 1},
		{-1, -2, -2},
		{-1, 2, -1},
		{1, -2, -2},
	}

	for _, tt := range test {
		testname := fmt.Sprintf("TestIntMinTableDriven(%d, %d)", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			r := IntMin(tt.a, tt.b)
			if r != tt.result {
				t.Errorf("get %d but expected %d", r, tt.result)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
