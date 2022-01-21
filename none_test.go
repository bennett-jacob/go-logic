package logic_test

import (
	"testing"

	"github.com/bennett-jacob/go-logic"
)

func TestNone(t *testing.T) {
	t1 := logic.None(false, false, false)
	t2 := logic.None(true, false, false)
	t3 := logic.None(false, true, true)
	t4 := logic.None(true, true, true)

	if t1 != true {
		t.Error("Xor(false, false, false) should be true")
	}
	if t2 != false {
		t.Error("Xor(true, false, false) should be false")
	}
	if t3 != false {
		t.Error("Xor(false, true, true) should be false")
	}
	if t4 != false {
		t.Error("Xor(true, true, true) should be false")
	}
}
