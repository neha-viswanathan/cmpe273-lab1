package fibo

import "testing"

type testValues struct {
	fiboIndex int
	fiboVal   int
}

var testFibo = []testValues{{1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {10, 100}}

func FiboTest(t *testing.T) {
	for _, pair := range testFibo {
		val := Fibo(pair.fiboIndex)
		if val != pair.fiboVal {
			t.Error("For ", pair.fiboIndex,
				"expected", pair.fiboVal,
				"got ", val,
			)
		}
	}
}
