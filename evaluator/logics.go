package evaluator

import (
	"fmt"
)

const (
	MIN_COUNT           = 2
	MIN_LETTER_SEQUENCE = 4
)

func (e *evaluator) Wait() bool {
	defer close(e.cantChan)

	count := 0
	func() {
		e.wg.Add(1)
		defer e.wg.Done()

		for {
			_, more := <-e.cantChan
			if more {
				count++
				fmt.Println("found one more: ", count)
				if count >= MIN_COUNT {
					e.safeStop()
					return
				}
			} else {
				fmt.Println("completed all the evaluation")
				return
			}
		}
	}()

	e.wg.Wait()

	return count >= MIN_COUNT
}

func (e *evaluator) Horizontal() {
	e.wg.Add(1)
	defer e.wg.Done()

	for _, row := range *e.dna {
		if e.shouldStop() {
			return
		}
		compare := e.charComparator()
		for key, _ := range row {
			compare(row[key])
		}
	}
}

func (e *evaluator) Vertical() {
	e.wg.Add(1)
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

func (e *evaluator) DiagonalRight() {
	e.wg.Add(1)
	defer e.wg.Done()

	length := len(*e.dna)
	compare := e.charComparator()
	for a := 0; a < length; a++ {
		compare((*e.dna)[a][a])
	}

	for a := 1; a < length-MIN_LETTER_SEQUENCE-1; a++ {
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

func (e *evaluator) DiagonalLeft() {
	e.wg.Add(1)
	defer e.wg.Done()

	length := len(*e.dna)
	corr := length-1
	compare := e.charComparator()
	for a := 0; a < length; a++ {
		compare((*e.dna)[a][corr-a])
	}

	for a := 1; a < length-MIN_LETTER_SEQUENCE-1; a++ {
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
