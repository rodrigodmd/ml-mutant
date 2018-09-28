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
	// testData{
	// 	testTitle:          "Null Dna",
	// 	dna:                []string{},
	// 	countHorizontal:    0,
	// 	countVertical:      0,
	// 	countDiagonalRight: 0,
	// 	countDiagonalLeft:  0,
	// },
	testData{
		testTitle:          "No Match Logic",
		dna:                []string{"ATGC", "CAGT", "TATT", "AGAA"},
		countHorizontal:    0,
		countVertical:      0,
		countDiagonalRight: 0,
		countDiagonalLeft:  0,
	},
	testData{
		testTitle:          "Horizontal Match",
		dna:                []string{"ATGCGA", "CAGTGC", "TTTTCT", "AGAAGG", "CCCCTA", "TCACTG"},
		countHorizontal:    2,
		countVertical:      0,
		countDiagonalRight: 0,
		countDiagonalLeft:  0,
	},
	testData{
		testTitle:          "Vertical Match",
		dna:                []string{"ATGCGA", "AAGTGC", "ATATGT", "AGACGG", "CACCTA", "TCACTG"},
		countHorizontal:    0,
		countVertical:      2,
		countDiagonalRight: 0,
		countDiagonalLeft:  0,
	},
	testData{
		testTitle:          "Horizontal And Vertical Match",
		dna:                []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
		countHorizontal:    1,
		countVertical:      1,
		countDiagonalRight: 1,
		countDiagonalLeft:  0,
	},
	testData{
		testTitle:          "Diagonal Leftl Match",
		dna:                []string{"ATGCGA", "CAGTAC", "TTAAGT", "AGAAGG", "CCCCTA", "TCACTG"},
		countHorizontal:    1,
		countVertical:      0,
		countDiagonalRight: 1,
		countDiagonalLeft:  1,
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
