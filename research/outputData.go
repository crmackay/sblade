package switchblade

import (
	//"encoding/csv"
	//"fmt"
	sw "github.com/crmackay/switchblade"
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

func newReadData(p *sw.OrigRead) readData {
	var numCuts, finalLen int
	var cutLens []int
	var avgPHREDBefore, avgPHREDAfter float64

	for _, elem := range p.ThreePTrims {
		if elem.IsLinker == true {
			cutLens = append(cutLens, len(elem.Alignment.Query)-elem.Alignment.QueryStart)
		}
	}
	numCuts = len(cutLens)

	if numCuts > 1 {
		finalLen = p.ThreePTrims[len(p.ThreePTrims)-2].Alignment.QueryStart
	} else {
		finalLen = len(p.Read.Sequence)
	}

	avgPHREDBefore = avgPHRED(p.Read.PHRED.Decoded)
	avgPHREDBefore = avgPHRED(p.Read.PHRED.Decoded[:finalLen])

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

func GetDataCSV(p *sw.InProcessRead) []string {
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
