package evaluator

import "sync"

func New(dna *[]string) *evaluator {
	return &evaluator{
		dna:            dna,
		chFound:        make(chan int),
		chProgress:     make(chan int),
		stop:           false,
	}

}

type evaluator struct {
	dna *[]string
	chFound        chan int
	chProgress     chan int
	stop           bool
	mutex          sync.Mutex
}
