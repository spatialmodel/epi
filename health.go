// Package epi holds a collection of functions for calculating the health impacts of air pollution.
package epi

import (
	"math"

	"github.com/gonum/floats"
)

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

// HR calculates the hazard ratio caused by concentration z.
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

// HRer is an interface for any type that can calculate the hazard ratio
// caused by concentration z.
type HRer interface {
	HR(z float64) float64
}

// IoRegional returns the underlying regional average incidence rate for a region where
// the reported incidence rate is I, individual locations within the
// region have population p and concentration z, and hr specifies the
// hazard ratio as a function of z, as presented in Equations 2 and 3 of:
//
// Apte JS, Marshall JD, Cohen AJ, Brauer M (2015) Addressing Global
// Mortality from Ambient PM2.5. Environmental Science and Technology
// 49(13):8057–8066.
func IoRegional(p, z []float64, hr HRer, I float64) float64 {
	var hrBar float64
	for i, pi := range p {
		hrBar += pi * hr.HR(z[i])
	}
	pSum := floats.Sum(p)
	hrBar /= pSum
	if pSum == 0 || hrBar == 0 {
		return 0
	}
	return I / hrBar
}

// Io returns the underlying incidence rate where
// the reported incidence rate is I, concentration is z,
// and hr specifies the hazard ratio as a function of z. When possible,
// IoRegional should be used instead of this function.
func Io(z float64, hr HRer, I float64) float64 {
	return I / hr.HR(z)
}

// Outcome returns the number of incidences occuring in population p when
// exposed to concentration z given underlying incidence rate Io and
// hazard relationship hr(z), as presented in Equation 2 of:
//
// Apte JS, Marshall JD, Cohen AJ, Brauer M (2015) Addressing Global
// Mortality from Ambient PM2.5. Environmental Science and Technology
// 49(13):8057–8066.
func Outcome(p, z, Io float64, hr HRer) float64 {
	return p * Io * (hr.HR(z) - 1)
}
