package evaluator

import "sync"

func New(dna *[]string) *evaluator {
	return &evaluator{
		dna:      dna,
		cantChan: make(chan int, 1),
		stop:     false,
	}

}

type evaluator struct {
	dna      *[]string
	wg       sync.WaitGroup
	cantChan chan int
	stop     bool
	mutex    sync.Mutex
}
