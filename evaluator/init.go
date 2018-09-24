package evaluator

import "sync"

func New(dna *[]string) *evaluator {
	e := evaluator{
		dna:      dna,
		count:    0,
		isMutant: false,
	}
	e.wg.Add(4)
	return &e
}

type evaluator struct {
	dna      *[]string
	count    int
	isMutant bool
	wg       sync.WaitGroup
	mutex    sync.Mutex
}
