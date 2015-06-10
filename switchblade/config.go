package switchblade

import (
//"fmt"
)

// default values
var pcrDetails = map[string]float64{
	"RTError":      0.0000003,
	"DNAPolError":  0.000000001,
	"NumPCRCycles": 20,
}

var pcrError = pcrDetails["RTError"] +
	(pcrDetails["DNAPolError"] * pcrDetails["NumPCRCycles"])
