package logic_test

import (
	"testing"

	"github.com/bennett-jacob/go-logic"
)

func TestAll(t *testing.T) {
	t1 := logic.All(false, false, false)
	t2 := logic.All(true, false, false)
	t3 := logic.All(false, true, true)
	t4 := logic.All(true, true, true)

	if t1 != false {
		t.Error("Xor(false, false, false) should be false")
	}
	if t2 != false {
		t.Error("Xor(true, false, false) should be false")
	}
	if t3 != false {
		t.Error("Xor(false, true, true) should be false")
	}
	if t4 != true {
		t.Error("Xor(true, true, true) should be true")
	}
}
