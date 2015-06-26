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
	read   sw.OrigRead
	result string
}

func newTestRead(r sw.OrigRead) testRead {
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
	testSuite := []testRead{newTestRead(sw.NewOrigRead(newRead))}
	for newRead.Sequence != nil {

		newTR := newTestRead(sw.NewOrigRead(newRead))

		testSuite = append(testSuite, newTR)
		newRead, _ = scanner.NextRead()

	}
	scanner.Close()
	return testSuite
}

// func next3pAlign(r sw.OrigRead) (bool, int) {
func TestNext3pAlign(t *testing.T) {
	fmt.Println("testing next3pAlign")
	//testSuite := setupTestData()
}

//func process3p(r sw.OrigRead) {
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
