package config

import (
	bio "github.com/crmackay/gobioinfo"
)

// default values
var PCRDetails = map[string]float64{
	"RTError":      0.00001490711984999862, // (http://www.chem.agilent.com/library/datasheets/public/tb108_71067.pdf)
	"DNAPolError":  0.00001038461538461538, // 26x taq https://www.neb.com/products/pcr-polymerases-and-amplification-technologies/q5-high-fidelity-dna-polymerases/q5-high-fidelity-dna-polymerases/how-is-fidelity-measured
	"NumPCRCycles": 30,
}

var PCRError = PCRDetails["RTError"] +
	(PCRDetails["DNAPolError"] * PCRDetails["NumPCRCycles"])

const ComplexityThreshold = 2

const LenFivePrimeLinker = 23

var Barcodes = map[string]map[string]string{
	"lane1": map[string]string{
		"AGC": "sample1",
		"CAC": "sample2",
		"GCT": "sample3",
		"GTG": "sample4",
	},
	"lane2": map[string]string{
		"AGC": "sample5",
		"CAC": "sample6",
		"GCT": "sample7",
		"GTG": "sample8",
	},
}

const Linker5p = "AGGGAGGACGATGCGGNNNNG"

var Linker3p = bio.NewDNASequence("GTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTG")
