package mlmutant

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/rodrigodmd/ml-mutant/evaluator"
)

func IsMutant(dna []string) (bool, error) {
	start := time.Now()

	if !checkStructure(&dna) {
		return false, errors.New("Invalid DNA structure")
	}

	evaluate := evaluator.New(&dna)
	go evaluate.Horizontal()
	go evaluate.Vertical()
	go evaluate.DiagonalLeft()
	go evaluate.DiagonalRight()

	result := evaluate.Wait()

	fmt.Println(time.Since(start))
	return result, nil
}

func checkStructure(dna *[]string) bool {
	length := len(*dna)
	if length == 0 {
		return false
	}

	for _, row := range *dna {
		if len(row) != length {
			return false
		}

		// TODO: Modify this logic to use regex
		// var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
		for _, letter := range row {
			if !strings.Contains("ATCG", string(letter)) {
				return false
			}
		}

	}
	return true
}
