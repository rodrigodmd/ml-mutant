package evaluator

const (
	minCount          = 2
	minLetterSequence = 4
)

// Wait method will wait until all the
// evaluation go routines finish the execution
// and returns the result
func (e *evaluator) Wait() bool {
	e.wg.Wait()
	return e.isMutant
}

// Horizontal evaluation process
func (e *evaluator) Horizontal() {
	defer e.wg.Done()

	for _, row := range *e.dna {
		if e.shouldStop() {
			return
		}
		compare := e.charComparator()
		for key := range row {
			compare(row[key])
		}
	}
}

// Vertical evaluation process
func (e *evaluator) Vertical() {
	defer e.wg.Done()

	length := len(*e.dna)
	for x := 0; x < length; x++ {
		if e.shouldStop() {
			return
		}
		compare := e.charComparator()
		for y := 0; y < length; y++ {
			compare((*e.dna)[y][x])
		}
	}

}

// Diagonal Right evaluation process
func (e *evaluator) DiagonalRight() {
	defer e.wg.Done()

	length := len(*e.dna)
	compare := e.charComparator()
	for a := 0; a < length; a++ {
		compare((*e.dna)[a][a])
	}

	for a := 1; a < length-minLetterSequence-1; a++ {
		if e.shouldStop() {
			return
		}
		compareUp := e.charComparator()
		compareDown := e.charComparator()
		for b := 0; a+b < length; b++ {
			compareUp((*e.dna)[b][a+b])
			compareDown((*e.dna)[a+b][b])
		}
	}

}

// Diagonal left evaluation process
func (e *evaluator) DiagonalLeft() {
	defer e.wg.Done()

	length := len(*e.dna)
	corr := length - 1
	compare := e.charComparator()
	for a := 0; a < length; a++ {
		compare((*e.dna)[a][corr-a])
	}

	for a := 1; a < length-minLetterSequence-1; a++ {
		if e.shouldStop() {
			return
		}
		compareUp := e.charComparator()
		compareDown := e.charComparator()
		for b := 0; a+b < length; b++ {
			compareUp((*e.dna)[a][corr-a-b])
			compareDown((*e.dna)[a+b][corr-a])
		}
	}
}
