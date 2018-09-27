package mlmutant

import (
	"errors"
	"regexp"
	"time"

	"github.com/rodrigodmd/ml-mutant/evaluator"
	"log"
)

// IsMutant method will verify the dna array structure
// Checking structure and finding sequence in the following
// logic:
// - Horizontal
// - Vertical
// - Diagonal right
// - Diagonal left
func IsMutant(dna []string) (bool, error) {
	start := time.Now()

	if !checkStructure(&dna) {
		return false, errors.New("Invalid DNA structure")
	}

	evaluate := evaluator.New(&dna, 4)
	go evaluate.Horizontal()
	go evaluate.Vertical()
	go evaluate.DiagonalLeft()
	go evaluate.DiagonalRight()

	result := evaluate.Wait()
	log.Print(time.Since(start))
	return result, nil
}

// checkStructure method checks the dna structure
// to se if it has the basic structure
func checkStructure(dna *[]string) bool {
	length := len(*dna)
	if length == 0 {
		return false
	}

	invalidDnaLetter := regexp.MustCompile(`[^ATCG]`).MatchString
	for _, row := range *dna {
		// Height should be equal to width for every row
		if len(row) != length {
			return false
		}

		// Chould contain only DNA letters
		if invalidDnaLetter(row) {
			return false
		}
	}
	return true
}
