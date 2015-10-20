package fibo

import "testing"

type fiboTest struct {
	in int
	out int
}

var testFibo = []fiboTest{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{100, 100},
}

func FiboTest(t *testing.T) {
	for _, pair := range testFibo {
		val := Fibo(pair.in)
		if val != pair.out{
			t.Error("For ", pair.in,
				"expected", pair.out,
				"got ", val,
			)
		}
	}
}
