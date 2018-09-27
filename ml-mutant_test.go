package mlmutant

import (
	"fmt"
	"testing"
)

type testData struct {
	testTitle         string
	dna               []string
	expectedIsMutant  bool
	shouldReturnError bool
}

var testTable = []testData{
	testData{
		testTitle:         "Null Dna",
		dna:               nil,
		shouldReturnError: true,
		expectedIsMutant:  false,
	},
	testData{
		testTitle:         "No Match Logic",
		dna:               []string{"ATGC", "CAGT", "TATT", "AGAA"},
		shouldReturnError: false,
		expectedIsMutant:  false,
	},
	testData{
		testTitle:         "Dna With Invalid Length",
		dna:               []string{"ATGCGA", "CAGTGCGCGC", "TTZZZZ", "AGAAGG", "CCCCTA", "TCACTG"},
		shouldReturnError: true,
		expectedIsMutant:  false,
	},
	testData{
		testTitle:         "Dna With Invalid Letters",
		dna:               []string{"ATGCGA", "CAGTGC", "TTZZZZ", "AGAAGG", "CCCCTA", "TCACTG"},
		shouldReturnError: true,
		expectedIsMutant:  false,
	},
	testData{
		testTitle:         "Horizontal Match",
		dna:               []string{"ATGCGA", "CAGTGC", "TTTTGT", "AGAAGG", "CCCCTA", "TCACTG"},
		shouldReturnError: false,
		expectedIsMutant:  true,
	},
	testData{
		testTitle:         "Vertical Match",
		dna:               []string{"ATGCGA", "AAGTGC", "ATATGT", "AGAAGG", "CACCTA", "TCACTG"},
		shouldReturnError: false,
		expectedIsMutant:  true,
	},
	testData{
		testTitle:         "Horizontal And Vertical Match",
		dna:               []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
		shouldReturnError: false,
		expectedIsMutant:  true,
	},
}

func TestIsMutantTable(t *testing.T) {
	for _, td := range testTable {
		result, err := IsMutant(td.dna)

		fmt.Printf("Testing %s \n", td.testTitle)
		if td.shouldReturnError && err == nil {
			t.Errorf("Should return an error for this scenario")
		}

		if !td.shouldReturnError && err != nil {
			fmt.Printf("ERROR: %v \n", err)
			t.Errorf("Should not return an error")
		}

		if td.expectedIsMutant != result {
			t.Errorf("isMutant should be %t", td.expectedIsMutant)
		}
	}
}
