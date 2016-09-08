package gotools

import (
	"reflect"
	"testing"
)

type AA struct {
	A string
	C int
}

type BB struct {
	AA
	D bool
	E map[string]string
}

type CC struct {
	A string
	C string
	D bool
	E string
}

func TestEmbedded(t *testing.T) {
	var cc CC
	bb := BB{
		AA: AA{
			A: "dsdd",
			C: 3,
		},
		D: true,
		E: map[string]string{"ab": "cc", "a": "ee"},
	}
	StructCopy(&cc, &bb)
	if cc.A != "dsdd" {
		t.Error("field A failed")
	}
	if cc.D != true {
		t.Error("field D failed")
	}
	if cc.C != "" {
		t.Error("field C failed")
	}
}

type WithPtr struct {
	A *string
	B int
	C *bool
	D *[]int
	E *string
}

type NoPtr struct {
	B int
	C bool
	D []int
	E string
	F bool
}

func TestSrcPtr(t *testing.T) {
	var a string = "aaa"
	var e string = "eeee"
	var c bool = true
	var d []int = []int{3, 2}
	src := WithPtr{
		A: &a,
		B: 3,
		C: &c,
		D: &d,
		E: &e,
	}

	var dst NoPtr
	StructCopy(&dst, &src)
	if src.B != dst.B {
		t.Error("field B failed")
	}
	if *src.C != dst.C {
		t.Error("field C failed")
	}
	if !reflect.DeepEqual(*src.D, dst.D) {
		t.Error("field D failed")
	}
	if *src.E != dst.E {
		t.Error("field E failed")
	}
}

func TestDstPtr(t *testing.T) {
	var e string = "eeee"
	var c bool = true
	var d []int = []int{3, 2}
	src := NoPtr{
		B: 3,
		C: c,
		D: d,
		E: e,
	}

	var dst WithPtr
	StructCopy(&dst, &src)
	if dst.A != nil {
		t.Error("field A failed")
	}
	if src.B != dst.B {
		t.Error("field B failed")
	}
	if *dst.C != src.C {
		t.Error("field C failed")
	}
	if !reflect.DeepEqual(*dst.D, src.D) {
		t.Error("field D failed")
	}
	if *dst.E != src.E {
		t.Error("field E failed")
	}
}
