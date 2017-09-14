// Package epi holds a collection of functions for calculating the health impacts of air pollution.
package epi

import "math"

// Nasari implements a class of simple approximations to the exposure response
// models described in:
//
// Nasari M, Szyszkowicz M, Chen H, Crouse D, Turner MC, Jerrett M, Pope CA III,
// Hubbell B, Fann N, Cohen A, Gapstur SM, Diver WR, Forouzanfar MH, Kim S-Y,
// Olives C, Krewski D, Burnett RT. (2015). A Class of Non-Linear
// Exposure-Response Models Suitable for Health Impact Assessment Applicable
// to Large Cohort Studies of Ambient Air Pollution.  Air Quality, Atmosphere,
// and Health: DOI: 10.1007/s11869-016-0398-z.
type Nasari struct {
	// Gamma, Delta, and Lambda are parameters fit using linear regression.
	Gamma, Delta, Lambda float64

	// F is the concentration transformation function.
	F func(z float64) float64
}

// HR calculates the health risk caused by concentration z.
func (n Nasari) HR(z float64) float64 {
	return math.Exp(n.Gamma * n.F(z) / (1 + math.Exp(-(z-n.Delta)/n.Lambda)))
}

// NasariACS is an exposure-response model fit to the American Cancer Society
// Cancer Prevention II cohort all causes of death from fine particulate matter.
var NasariACS = Nasari{
	Gamma:  0.0478,
	Delta:  6.94,
	Lambda: 3.37,
	F:      func(z float64) float64 { return math.Log(z + 1) },
}
