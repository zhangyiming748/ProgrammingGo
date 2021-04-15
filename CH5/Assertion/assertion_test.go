package Assertion

import "testing"

func TestAssertion(t *testing.T) {
	var (
		float  float32 = 1.2
		double float64 = 1.2
	)

	t.Log(assertion(float))
	t.Log(assertion(double))
}
