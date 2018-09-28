package evaluator

import "log"

// shouldStop method verifies the flag to stop
// all running evaluation process
func (e *evaluator) shouldStop() bool {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	return e.isMutant
}

// Each time we find a sequence, we increase the counter
// if the counter is greater than the threshold for mutants,
// we change the flag to stop all process
func (e *evaluator) foundSequence() {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	if e.count++; e.count == minCount {
		e.isMutant = true
	}
}

// charComparator creates a function that compares
// the current letter with the last letter and counts
func (e *evaluator) charComparator() func(uint8) {
	var lastChar uint8 = ' '
	count := 1

	log.Printf("Restart count")
	return func(char uint8) {
		if lastChar == char {
			count++
		} else {
			count = 1
		}
		log.Printf("char "+string(char)+"  count: %d", count)
		if count == minLetterSequence {
			e.foundSequence()
		}
		lastChar = char
	}
}
