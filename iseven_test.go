package iseven

import "testing"

type isevenTest struct {
	number int
	parity bool
}

var isevenTests = []isevenTest{
	{0, true},
	{-4, true},
	{33, false},
	{-69, false},
	{100, true},
	{-2137, false},
	{100000, true},
	{111110, true},
	{312321, false},
	{5009123, false},
}

func TestIsEven(t *testing.T) {
	for _, test := range isevenTests {
		if parity := IsEven(test.number); parity != test.parity {
			t.Errorf("Number's %d output is %t, but %t expected", test.number, parity, test.parity)
		}
	}
}
