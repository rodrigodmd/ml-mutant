package evaluator



func (e *evaluator) shouldStop() bool {
	return e.stop
}

func (e *evaluator) safeStop() {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.stop = true
}

func (e *evaluator) safeAddWaitGroup() {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.wg.Add(1)
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

		//log.Print(lastChar, " vs " , char, "    Count: " , count)
		if count == MIN_LETTER_SEQUENCE {
			e.cantChan <- 1
		}
		lastChar = char
	}
}
