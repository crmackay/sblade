package switchblade

import (
	"fmt"
	bio "github.com/crmackay/gobioinfo"
	"testing"
)

func TestLinkerTest(t *testing.T) {

	//SET GLOBAL PCR VARIABLES for testing

	PCRDetails := make(map[string]float64)

	PCRDetails["RTError"] = 0.0000003

	PCRDetails["DNAPolError"] = 0.000000001

	PCRDetails["NumPCRCycles"] = 20

	// create test alignment struct

	testAlignment := bio.PairWiseAlignment{
		Subject:                 bio.NucleotideSequence([]rune("GTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTGCTTG")),
		Query:                   bio.NucleotideSequence([]rune("GCTAGGGAGGACGATGCGGTGGTGATGCTGCCACATACACTAAGAAGGTCCTGGACGCGTGTAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAA")),
		ExpandedCIGAR:           []string{"m", "m", "m", "m", "i", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "j", "m", "j", "m", "m", "m", "m", "m", "m"},
		SubjectStart:            0,
		QueryStart:              58,
		SubjectAlignLen:         40,
		QueryAlignLen:           41,
		GappedSubject:           "GTGTCAGTCACTTCCAGCGGTCGTATGCCGTC-T-TGCTTG",
		GappedQuery:             "GTGT-AGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTG",
		AlignmentRepresentation: "|||| ||||||||||||||||||||||||||| | ||||||"}

	// TODO Create a FASTQ Read (use the bioinfo package and four strings to do this)

	testQuery := bio.FASTQRead{
		Id: "test query",
		DNASequence: bio.DNASequence{
			Sequence: testAlignment.Query,
		},
		Misc: "",
		QSequence: bio.QSequence{
			QualByteSequence: []rune("@@@FFFFFHHFFFFFHGHJ@FH?BFHF<HIGGIJIGJJGG=CCGGGHIC@=DDECHHED3>@CDCDCACC>>@A:9>99@)<>?@>@5)8<@CC:A>A<A"),
			PHRED:            bio.DecodeQualByteSequence([]rune("@@@FFFFFHHFFFFFHGHJ@FH?BFHF<HIGGIJIGJJGG=CCGGGHIC@=DDECHHED3>@CDCDCACC>>@A:9>99@)<>?@>@5)8<@CC:A>A<A"), "Illumina 1.8"),
			Encoding:         "Illumina 1.8",
		},
	}

	testSubject := testAlignment.Query

	testLinkerTestSet := linkerTestSet{
		testAlignment,
		testQuery,
		testSubject,
		PCRDetails,
	}

	result := LinkerTest(testLinkerTestSet)

	fmt.Println(result)

	if result != true {
		t.Error("testLinker returned 'false' instead of 'true'")
	}

	fmt.Println(result)
}

func TestTrimRead(t *testing.T) {

	// create series of (alignment, bool, and trimmed alignments)
	// trimmedRead := TrimRead(in bio.PairWiseAlignment, hasLinker bool)
	// test trimmedRead

}

func TestCLI(t *testing.T) {
}
