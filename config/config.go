package config

// default values
var PCRDetails = map[string]float64{
	"RTError":      0.00001490711984999862, // (http://www.chem.agilent.com/library/datasheets/public/tb108_71067.pdf)
	"DNAPolError":  0.00001038461538461538, // 26x taq https://www.neb.com/products/pcr-polymerases-and-amplification-technologies/q5-high-fidelity-dna-polymerases/q5-high-fidelity-dna-polymerases/how-is-fidelity-measured
	"NumPCRCycles": 30,
}

var PCRError = PCRDetails["RTError"] +
	(PCRDetails["DNAPolError"] * PCRDetails["NumPCRCycles"])

const ComplexityThreshold = 2
