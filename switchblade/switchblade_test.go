package switchblade

import (
	"fmt"
	bio "github.com/crmackay/gobioinfo"
	"testing"
)

func TestLinkerTest(t *testing.T) {

	//SET GLOBAL PCR VARIABLES for testing

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
		ID: "test query",
		DNASequence: bio.DNASequence{
			Sequence: testAlignment.Query,
		},
		Misc: "",
		PHRED: bio.PHRED{
			Encoded:  []rune("@@@FFFFFHHFFFFFHGHJ@FH?BFHF<HIGGIJIGJJGG=CCGGGHIC@=DDECHHED3>@CDCDCACC>>@A:9>99@)<>?@>@5)8<@CC:A>A<A"),
			Decoded:  nil,
			Encoding: "illumina_1.8",
		},
	}
	testQuery.PHRED.Decode()

	testSubject := bio.DNASequence{Sequence: testAlignment.Subject}

	testLinkerTestSet := alignmentSet{
		testAlignment,
		testQuery,
		testSubject,
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
