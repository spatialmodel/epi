package epi

import (
	"fmt"
	"testing"
)

func TestNasariACS(t *testing.T) {
	var tests = []struct {
		in, out float64
	}{
		{
			in:  0,
			out: 1,
		},
		{
			in:  5,
			out: 1.031306668121412,
		},
		{
			in:  15,
			out: 1.1291019999220953,
		},
		{
			in:  25,
			out: 1.1676668889134683,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.in), func(t *testing.T) {
			have := NasariACS.HR(test.in)
			if have != test.out {
				t.Errorf("%g = %g, want %g", test.in, have, test.out)
			}
		})
	}

}
