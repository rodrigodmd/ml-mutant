package evaluator

import "log"

func (e *evaluator) shouldStop() bool {
	return e.isMutant
}

func (e *evaluator) foundSequence() {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	if e.count++; e.count == MIN_COUNT {
		e.isMutant = true
	}
	log.Print(e.count)
}

func (e *evaluator) charComparator() func(uint8) {
	var lastChar uint8 = ' '
	count := 1

	return func(char uint8) {
		if lastChar == char {
			count++
		} else {
			count = 1
		}

		//log.Print(lastChar, " vs ", char, "    Count: ", count)
		if count == MIN_LETTER_SEQUENCE {
			log.Print("FOUND")
			e.foundSequence()
		}
		lastChar = char
	}
}
