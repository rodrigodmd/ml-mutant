package evaluator

import "sync"

// New method returns an evaluator instance.
// It contains all the DNA evaluation logic:
// - Horizontal
// - Vertical
// - Diagonal right
// - Diagonal left
//
// Since evaluation is done in parallel, we need
// evaluatorCount to wait for that amount of go routines to finish
func New(dna *[]string, evaluatorCount int) *evaluator {
	e := evaluator{
		dna:      dna,
		count:    0,
		isMutant: false,
	}
	e.wg.Add(evaluatorCount)
	return &e
}

// evaluator basic structure
type evaluator struct {
	dna      *[]string
	count    int
	isMutant bool
	wg       sync.WaitGroup
	mutex    sync.Mutex
}
