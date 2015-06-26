package threeprime

import (
	"bufio"
	"bytes"
	"fmt"
	bio "github.com/crmackay/gobioinfo"
	sw "github.com/crmackay/switchblade"
	"testing"
)

type testRead struct {
	read   sw.Read
	result string
}

func newTestRead(r sw.Read) testRead {
	newTR := testRead{
		read:   r,
		result: r.Misc[1:],
	}
	return newTR
}

func setupTestData() []testRead {
	testData := bytes.NewBufferString(testRawReads)
	scanner := bio.FASTQScanner{Scanner: bufio.NewScanner(testData), File: nil}
	newRead, _ := scanner.NextRead()
	testSuite := []testRead{newTestRead(sw.NewRead(newRead))}
	for newRead.Sequence != nil {

		newTR := newTestRead(sw.NewRead(newRead))

		testSuite = append(testSuite, newTR)
		newRead, _ = scanner.NextRead()

	}
	scanner.Close()
	return testSuite
}

// func next3pAlign(r sw.Read) (bool, int) {
func TestNext3pAlign(t *testing.T) {
	fmt.Println("testing next3pAlign")
	//testSuite := setupTestData()
}

//func process3p(r sw.Read) {
func TestProcess3p(t *testing.T) {
	fmt.Println("testing process3p()")
	testSuite := setupTestData()
	for _, test := range testSuite {
		process3p(&test.read)
		if string(test.read.FinalSeq) != test.result {
			t.Error("expecting:\t",
				test.result,
				"\n\tbut got:\t",
				string(test.read.FinalSeq))
		}
	}
}
