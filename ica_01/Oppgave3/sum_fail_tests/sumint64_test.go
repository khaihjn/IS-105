package sum

import (
	"testing"

	sum "github.com/SondreH/IS-105/ica_01/Oppgave3/sum"
)

var sum_tests_int64 = []struct {
	n1       int64
	n2       int64
	expected int64
}{
	{1, 2, 3},
	{4, -5, -1},
	{12657999996666897987, 1, 127},
}

func TestSumInt64(t *testing.T) {
	for _, v := range sum_tests_int64 {
		if val := sum.SumInt64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}
