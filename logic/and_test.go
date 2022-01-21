package logic_test

import (
	"testing"

	"github.com/bennett-jacob/go-logic/logic"
)

func TestAnd(t *testing.T) {
	t1 := logic.And(false, false, false)
	t2 := logic.And(true, false, false)
	t3 := logic.And(false, true, true)
	t4 := logic.And(true, true, true)

	if t1 != false {
		t.Error("And(false, false, false) should be false")
	}
	if t2 != false {
		t.Error("And(true, false, false) should be false")
	}
	if t3 != false {
		t.Error("And(false, true, true) should be false")
	}
	if t4 != true {
		t.Error("And(true, true, true) should be true")
	}
}

func TestAll(t *testing.T) {
	t1 := logic.All(false, false, false)
	t2 := logic.All(true, false, false)
	t3 := logic.All(false, true, true)
	t4 := logic.All(true, true, true)

	if t1 != false {
		t.Error("All(false, false, false) should be false")
	}
	if t2 != false {
		t.Error("All(true, false, false) should be false")
	}
	if t3 != false {
		t.Error("All(false, true, true) should be false")
	}
	if t4 != true {
		t.Error("All(true, true, true) should be true")
	}
}
