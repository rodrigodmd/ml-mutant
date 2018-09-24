package evaluator

import (
	"fmt"
)

const (
	MIN_COUNT           = 2
	MIN_LETTER_SEQUENCE = 4
)

func (e *evaluator) Wait() bool {
	e.wg.Wait()
	return e.isMutant
}

func (e *evaluator) Horizontal() {
	defer e.wg.Done()
	defer fmt.Println("DONE Horizontal")

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
	defer e.wg.Done()
	defer fmt.Println("DONE Vertical")

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
	defer e.wg.Done()
	defer fmt.Println("DONE DiagonalRight")

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
	defer e.wg.Done()
	defer fmt.Println("DONE DiagonalLeft")

	length := len(*e.dna)
	corr := length - 1
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
