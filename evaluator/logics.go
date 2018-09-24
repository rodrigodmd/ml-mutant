package evaluator

import (
	"fmt"
	"sync"
)

const (
	MIN_COUNT           = 2
	MIN_LETTER_SEQUENCE = 4
)

func (e *evaluator) Wait() bool {
	var wg sync.WaitGroup
	wg.Add(1)

	isMutant := false
	//func (){
	//defer wg.Done()
	foundCount, progressCount := 0, 0
	for{
		select {
		case <-e.chFound:
			foundCount++
			fmt.Println("got found")
			if foundCount == MIN_COUNT{
				e.safeStop()
				isMutant = true
				return true
			}
		case delta := <-e.chProgress:
			if progressCount+=delta; progressCount == 0{
				isMutant = false
				return false
			}
		}
	}
	//}()

	//wg.Wait()
	return isMutant
}

func (e *evaluator) Horizontal() {
	e.chProgress <- 1

	for _, row := range *e.dna {
		if e.shouldStop() {
			return
		}
		compare := e.charComparator()
		for key, _ := range row {
			compare(row[key])
		}
	}
	fmt.Println("DONE Horizontal")
	e.chProgress <- -1
}

func (e *evaluator) Vertical() {
	e.chProgress <- 1

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
	fmt.Println("DONE Vertical")
	e.chProgress <- -1
}

func (e *evaluator) DiagonalRight() {
	e.chProgress <- 1

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
	fmt.Println("DONE DiagonalRight")
	e.chProgress <- -1
}

func (e *evaluator) DiagonalLeft() {
	e.chProgress <- 1

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
	fmt.Println("DONE DiagonalLeft")
	e.chProgress <- -1
}
