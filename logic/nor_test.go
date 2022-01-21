package logic_test

import (
	"testing"

	"github.com/bennett-jacob/go-logic/logic"
)

func TestNor(t *testing.T) {
	t1 := logic.Nor(false, false, false)
	t2 := logic.Nor(true, false, false)
	t3 := logic.Nor(false, true, true)
	t4 := logic.Nor(true, true, true)

	if t1 != true {
		t.Error("Nor(false, false, false) should be true")
	}
	if t2 != false {
		t.Error("Nor(true, false, false) should be false")
	}
	if t3 != false {
		t.Error("Nor(false, true, true) should be false")
	}
	if t4 != false {
		t.Error("Nor(true, true, true) should be false")
	}
}

func TestNone(t *testing.T) {
	t1 := logic.None(false, false, false)
	t2 := logic.None(true, false, false)
	t3 := logic.None(false, true, true)
	t4 := logic.None(true, true, true)

	if t1 != true {
		t.Error("None(false, false, false) should be true")
	}
	if t2 != false {
		t.Error("None(true, false, false) should be false")
	}
	if t3 != false {
		t.Error("None(false, true, true) should be false")
	}
	if t4 != false {
		t.Error("None(true, true, true) should be false")
	}
}
