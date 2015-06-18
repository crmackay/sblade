package config

// default values
var PcrDetails = map[string]float64{
	"RTError":      0.0001490711984999862, // (http://www.chem.agilent.com/library/datasheets/public/tb108_71067.pdf)
	"DNAPolError":  0.000004230769230769231,
	"NumPCRCycles": 30,
}

var PcrError = PcrDetails["RTError"] +
	(PcrDetails["DNAPolError"] * PcrDetails["NumPCRCycles"])
