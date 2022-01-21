package logic_test

import (
	"testing"

	"github.com/bennett-jacob/go-logic/logic"
)

func TestOnlyOne(t *testing.T) {
	t1 := logic.OnlyOne(false, false)
	t2 := logic.OnlyOne(true, false)
	t3 := logic.OnlyOne(false, true)
	t4 := logic.OnlyOne(true, true)

	if t1 != false {
		t.Error("OnlyOne(false, false) should be false")
	}
	if t2 != true {
		t.Error("OnlyOne(true, false) should be true")
	}
	if t3 != true {
		t.Error("OnlyOne(false, true) should be true")
	}
	if t4 != false {
		t.Error("OnlyOne(true, true) should be false")
	}
}

func TestXor(t *testing.T) {
	t1 := logic.Xor(false, false)
	t2 := logic.Xor(true, false)
	t3 := logic.Xor(false, true)
	t4 := logic.Xor(true, true)

	if t1 != false {
		t.Error("Xor(false, false) should be false")
	}
	if t2 != true {
		t.Error("Xor(true, false) should be true")
	}
	if t3 != true {
		t.Error("Xor(false, true) should be true")
	}
	if t4 != false {
		t.Error("Xor(true, true) should be false")
	}
}
