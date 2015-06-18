package switchblade

import (
	//"encoding/csv"
	//"fmt"
	"strconv"
	"strings"
)

/* fields that we want to collect per read:

- number of cuts conducted
- final length
- length of each cut
- avg PHRED score before and after

fields that we want to collect per file:
- total number of reads
- total number of linkers removed
- avg final read length
- avg before PHRED score
- avg final PHRED score

*/

type readData struct {
	numCuts        int
	finalLen       int
	cutLens        []int
	avgPHREDBefore float64
	avgPHREDAfter  float64
}

type fileData struct {
	numReads    int
	numLinkers  int
	avgFinalLen float64
}

func newReadData(p *inProcessRead) readData {
	var numCuts, finalLen int
	var cutLens []int
	var avgPHREDBefore, avgPHREDAfter float64

	for _, elem := range p.threePTrims {
		if elem.isLinker == true {
			cutLens = append(cutLens, len(elem.alignment.Query)-elem.alignment.QueryStart)
		}
	}
	numCuts = len(cutLens)

	if numCuts > 1 {
		finalLen = p.threePTrims[len(p.threePTrims)-2].alignment.QueryStart
	} else {
		finalLen = len(p.read.Sequence)
	}

	avgPHREDBefore = avgPHRED(p.read.PHRED.Decoded)
	avgPHREDBefore = avgPHRED(p.read.PHRED.Decoded[:finalLen])

	newData := readData{
		numCuts:        numCuts,
		finalLen:       finalLen,
		cutLens:        cutLens,
		avgPHREDBefore: avgPHREDBefore,
		avgPHREDAfter:  avgPHREDAfter,
	}
	return newData
}

func avgPHRED(s []uint8) float64 {
	var sum int
	for _, elem := range s {
		sum += int(elem)
	}
	avg := float64(sum) / float64(len(s))
	return avg
}

func (p *inProcessRead) getDataCSV() []string {
	newData := newReadData(p)

	toCSV := []string{
		strconv.Itoa(newData.numCuts),
		strconv.Itoa(newData.finalLen),
		arrayToString(newData.cutLens),
		strconv.FormatFloat(newData.avgPHREDBefore, 'f', 5, 64),
		strconv.FormatFloat(newData.avgPHREDAfter, 'f', 5, 64),
	}

	return toCSV

}

func arrayToString(a []int) (s string) {

	newStrings := make([]string, len(a))

	for i := range a {
		newStrings[i] = strconv.Itoa(int(a[i]))
	}

	s = strings.Join(newStrings, ",")
	return s
}
