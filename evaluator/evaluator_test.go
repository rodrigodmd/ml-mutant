package evaluator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTest() {

}

type testData struct {
	testTitle          string
	dna                []string
	countHorizontal    int
	countVertical      int
	countDiagonalRight int
	countDiagonalLeft  int
}

var testTable = []testData{
	testData{
		testTitle:          "Null Dna",
		dna:                []string{},
		countHorizontal:    0,
		countVertical:      0,
		countDiagonalRight: 0,
		countDiagonalLeft:  0,
	},
	testData{
		testTitle:          "No Match Logic",
		dna:                []string{"ABCDEF", "GHIJKL", "MNOPQR", "ABCDEF", "GHIJKL", "MNOPQR"},
		countHorizontal:    0,
		countVertical:      0,
		countDiagonalRight: 0,
		countDiagonalLeft:  0,
	},
	testData{
		testTitle:          "Horizontal Match",
		dna:                []string{"ABCDEF", "GHIJKL", "MN2222", "ABCDEF", "GHIJKL", "M3333R"},
		countHorizontal:    2,
		countVertical:      0,
		countDiagonalRight: 0,
		countDiagonalLeft:  0,
	},
	testData{
		testTitle:          "Vertical Match",
		dna:                []string{"1BCDEF", "1HIJKL", "1NOPQ2", "1BCDE2", "GHIJK2", "MNOPQ2"},
		countHorizontal:    0,
		countVertical:      2,
		countDiagonalRight: 0,
		countDiagonalLeft:  0,
	},
	testData{
		testTitle: "Horizontal And Vertical Match",
		dna: []string{
			"1BCDEF",
			"1HIJK2",
			"1NOPQ2",
			"1111E2",
			"GHIJK2",
			"MN3333"},
		countHorizontal:    2,
		countVertical:      2,
		countDiagonalRight: 0,
		countDiagonalLeft:  0,
	},
	testData{
		testTitle: "Diagonal Leftl Match",
		dna: []string{
			"AB1DEF",
			"GHI1KL",
			"MN1P1R",
			"ABC1E1",
			"GHIJ1L",
			"MNOPQ1"},
		countHorizontal:    0,
		countVertical:      0,
		countDiagonalRight: 2,
		countDiagonalLeft:  0,
	},
	testData{
		testTitle: "Diagonal Leftl Bottom Match",
		dna: []string{
			"ABCDEF",
			"GHIJKL",
			"M1OPQR",
			"AB1DEF",
			"GHI1KL",
			"MNOP1R"},
		countHorizontal:    0,
		countVertical:      0,
		countDiagonalRight: 1,
		countDiagonalLeft:  0,
	},

	testData{
		testTitle: "Diagonal Righe Match",
		dna: []string{
			"ABCDE1",
			"GHIJ1L",
			"MNO1Q2",
			"AB1D2F",
			"GHI2KL",
			"MN2PQR"},
		countHorizontal:    0,
		countVertical:      0,
		countDiagonalRight: 0,
		countDiagonalLeft:  2,
	},

	testData{
		testTitle: "Diagonal Righe Match",
		dna: []string{
			"ABCD1F",
			"GHI1KL",
			"MN1PQ2",
			"A1CD2F",
			"GHI2KL",
			"MN2PQR"},
		countHorizontal:    0,
		countVertical:      0,
		countDiagonalRight: 0,
		countDiagonalLeft:  2,
	},
}

func TestTableHorizontal(t *testing.T) {
	assert := assert.New(t)
	for _, td := range testTable {
		e := New(&td.dna, 1)
		go e.Horizontal()
		result := e.Wait()

		fmt.Printf("Testing %s \n", td.testTitle)
		assert.Equal(td.countHorizontal, e.count, "Horizontal Count should be %d", td.countHorizontal)
		assert.Equal(td.countHorizontal > 1, result, "Result should be %v", td.countHorizontal > 1)
	}
}

func TestTableVertical(t *testing.T) {
	assert := assert.New(t)
	for _, td := range testTable {

		fmt.Printf("Testing %s \n", td.testTitle)
		e := New(&td.dna, 1)
		go e.Vertical()
		result := e.Wait()
		assert.Equal(td.countVertical, e.count, "Horizontal Count should be %d", td.countVertical)
		assert.Equal(td.countVertical > 1, result, "Result should be %v", td.countVertical > 1)
	}
}

func TestTableDiagonalRight(t *testing.T) {
	assert := assert.New(t)
	for _, td := range testTable {

		fmt.Printf("Testing %s \n", td.testTitle)
		e := New(&td.dna, 1)
		go e.DiagonalRight()
		result := e.Wait()
		assert.Equal(td.countDiagonalRight, e.count, "Horizontal Count should be %d", td.countDiagonalRight)
		assert.Equal(td.countDiagonalRight > 1, result, "Result should be %v", td.countDiagonalRight > 1)
	}
}

func TestTableDiagonalLeft(t *testing.T) {
	assert := assert.New(t)
	for _, td := range testTable {

		fmt.Printf("Testing %s \n", td.testTitle)
		e := New(&td.dna, 1)
		go e.DiagonalLeft()
		result := e.Wait()
		assert.Equal(td.countDiagonalLeft, e.count, "Horizontal Count should be %d", td.countDiagonalLeft)
		assert.Equal(td.countDiagonalLeft > 1, result, "Result should be %v", td.countDiagonalLeft > 1)
	}
}
