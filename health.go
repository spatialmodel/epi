package aqhealth

import "math"

// Relative risk from PM2.5 concentration change, assuming a
// log-log dose response (almost a linear relationship).
// From From Krewski et al (2009, Table 11)
// and Josh Apte (personal communication).
func RRpm25Linear(pm25 float64) float64 {
	return math.Exp(PM25linearCoefficient * pm25)
}

var PM25linearCoefficient = 0.007510747

// Relative risk from PM2.5 concentration change, assuming a
// log-linear dose response. From From Krewski et al (2009, Table 11)
// and Josh Apte (personal communication).
func RRpm25Log(baselinePM25, Δpm25 float64) float64 {
	var RR = 0.
	if baselinePM25 != 0. {
		RR = PM25logCoefficient * math.Pow((Δpm25+baselinePM25)/
			baselinePM25, 0.109532154)
	}
	return RR
}

var PM25logCoefficient = 1.000112789
var PM25logExponent = 0.109532154

// Relative risk from O3 concentration change, assuming a
// log-log dose response (almost a linear relationship).
// From From Jerrett et al (2009)
// and Josh Apte (personal communication).
func RRo3Linear(o3 float64) float64 {
	return math.Exp(O3linearCoefficient * o3)
}

var O3linearCoefficient = 0.003922071

// MR (baseline mortality rate) is in units of deaths per 100,000 people
// per year.
// people * deathsPer100,000 / 100,000 * RR = delta deaths
func Deaths(RR, population, MR float64) float64 {
	return (RR - 1) * population * MR / 100000.
}
