package evaluator

import "log"

func (e *evaluator) shouldStop() bool {
	return e.stop
}

func (e *evaluator) safeStop() {
	//e.mutex.Lock()
	//defer e.mutex.Unlock()
	e.stop = true
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
			e.chFound <- 1
		}
		lastChar = char
	}
}
