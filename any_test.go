package logic_test

import (
	"testing"

	"github.com/bennett-jacob/go-logic"
)

func TestAny(t *testing.T) {
	t1 := logic.Any(false, false, false)
	t2 := logic.Any(true, false, false)
	t3 := logic.Any(false, true, true)
	t4 := logic.Any(true, true, true)

	if t1 != false {
		t.Error("Xor(false, false, false) should be false")
	}
	if t2 != true {
		t.Error("Xor(true, false, false) should be true")
	}
	if t3 != true {
		t.Error("Xor(false, true, true) should be true")
	}
	if t4 != true {
		t.Error("Xor(true, true, true) should be true")
	}
}
