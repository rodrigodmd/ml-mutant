package evaluator

import (
	"testing"
)

func setupTest() {

}

func TestAa(t *testing.T) {

	e := New(&[]string{}, 1)
	go e.Horizontal()
	result := e.Wait()
	if result {

	}
}
