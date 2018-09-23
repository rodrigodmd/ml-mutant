package mlMutant

import (
	"testing"

	"log"
)

func TestNullDna(t *testing.T) {

	res, err := IsMutant(nil)
	if err != nil {
		log.Print(err)
	} else {
		t.Errorf("Should return an error for empty dna")
	}

	if res {
		t.Errorf("isMutant should return true")
	}

	log.Print("Result: ", res)
}

func TestDnaWithInvalidLetters(t *testing.T) {

	dna := []string{"ATGCGA","CAGTGC","TTZZZZ","AGAAGG","CCCCTA","TCACTG"};
	res, err := IsMutant(dna)

	if err != nil {
		log.Print(err)
	} else {
		t.Errorf("Should return an error for empty dna")
	}

	if res {
		t.Errorf("isMutant should return true")
	}

	log.Print("Result: ", res)
}

func TestHorizontalLogic(t *testing.T) {
	dna := []string{"ATGCGA","CAGTGC","TTTTGT","AGAAGG","CCCCTA","TCACTG"};
	res, err := IsMutant(dna)

	if err != nil {
		t.Error(err)
	}

	if !res {
		t.Errorf("isMutant should return true")
	}

	log.Print("Result: ", res)
}

func TestVerticalLogic(t *testing.T) {
	dna := []string{"ATGCGA","AAGTGC","ATATGT","AGAAGG","CACCTA","TCACTG"};
	res, err := IsMutant(dna)

	if err != nil {
		t.Error(err)
	}

	if !res {
		t.Errorf("isMutant should return true")
	}

	log.Print("Result: ", res)
}

func TestHorizontalAndVerticalLogic(t *testing.T) {
	dna := []string{"ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"};
	res, err := IsMutant(dna)

	if err != nil {
		t.Error(err)
	}

	if !res {
		t.Errorf("isMutant should return true")
	}

	log.Print("Result: ", res)
}