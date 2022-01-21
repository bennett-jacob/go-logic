package logic_test

import (
	"testing"

	"github.com/bennett-jacob/go-logic/logic"
)

func TestOr(t *testing.T) {
	t1 := logic.Or(false, false, false)
	t2 := logic.Or(true, false, false)
	t3 := logic.Or(false, true, true)
	t4 := logic.Or(true, true, true)

	if t1 != false {
		t.Error("Or(false, false, false) should be false")
	}
	if t2 != true {
		t.Error("Or(true, false, false) should be true")
	}
	if t3 != true {
		t.Error("Or(false, true, true) should be true")
	}
	if t4 != true {
		t.Error("Or(true, true, true) should be true")
	}
}

func TestAny(t *testing.T) {
	t1 := logic.Any(false, false, false)
	t2 := logic.Any(true, false, false)
	t3 := logic.Any(false, true, true)
	t4 := logic.Any(true, true, true)

	if t1 != false {
		t.Error("Any(false, false, false) should be false")
	}
	if t2 != true {
		t.Error("Any(true, false, false) should be true")
	}
	if t3 != true {
		t.Error("Any(false, true, true) should be true")
	}
	if t4 != true {
		t.Error("Any(true, true, true) should be true")
	}
}
