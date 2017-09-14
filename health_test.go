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

// This example calculates mortalities caused by ambient PM2.5 concentrations.
func Example() {
	var (
		// I represents currently observed deaths per 100,000
		// people in this region.
		I = 800.0

		// p represents number of people in different locations
		// in the region.
		p = []float64{100000, 80000, 700000, 90000}

		// z represents PM2.5 concentrations in μg/m³ in locations
		// corresponding to p.
		z = []float64{12, 26, 11, 2, 9}
	)

	// This is how we can calculate total deaths caused by air pollution
	// using a regional underlying incidence rate.
	io := IoRegional(p, z, NasariACS, I/100000)
	var totalDeaths float64
	for i, pi := range p {
		totalDeaths += Outcome(pi, z[i], io, NasariACS)
	}
	fmt.Printf("total deaths using regional underlying incidence: %.0f\n", totalDeaths)

	// This is how we can calculate additional deaths caused by doubling
	// air pollution:
	var doubleDeaths float64
	for i, pi := range p {
		doubleDeaths += Outcome(pi, z[i]*2, io, NasariACS) - Outcome(pi, z[i], io, NasariACS)
	}
	fmt.Printf("additional deaths caused by doubling air pollution: %.0f\n", doubleDeaths)

	// This is how we can calculate lives saved by halving air pollution:
	var halfDeaths float64
	for i, pi := range p {
		halfDeaths += Outcome(pi, z[i]/2, io, NasariACS) - Outcome(pi, z[i], io, NasariACS)
	}
	fmt.Printf("lives saved by halving air pollution: %.0f\n", -1*halfDeaths)

	// Sometimes it is not practical to calculate regional underlying
	// incidence. This is how we can calculate total deaths caused by air pollution
	// using a location-specific underlying incidence rate.
	totalDeaths = 0
	for i, pi := range p {
		totalDeaths += Outcome(pi, z[i], Io(z[i], NasariACS, I/100000), NasariACS)
	}
	fmt.Printf("total deaths using local underlying incidence: %.0f\n", totalDeaths)

	// This is how we can calculate additional deaths caused by doubling
	// air pollution using a local underlying incidence rate:
	doubleDeaths = 0
	for i, pi := range p {
		io := Io(z[i], NasariACS, I/100000)
		doubleDeaths += Outcome(pi, z[i]*2, io, NasariACS) - Outcome(pi, z[i], io, NasariACS)
	}
	fmt.Printf("additional deaths caused by doubling air pollution (local underlying incidence): %.0f\n", doubleDeaths)

	// Output:
	// total deaths using regional underlying incidence: 672
	// additional deaths caused by doubling air pollution: 403
	// lives saved by halving air pollution: 389
	// total deaths using local underlying incidence: 665
	// additional deaths caused by doubling air pollution (local underlying incidence): 401
}
