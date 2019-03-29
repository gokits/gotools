package gotools

import (
	"testing"
)

func TestNewSampler(t *testing.T) {
	var err error
	type testNewParam struct {
		ratio   float32
		width   uint32
		success bool
	}
	params := []testNewParam{
		testNewParam{-1, 1000, false},
		testNewParam{0, 1000, true},
		testNewParam{0.3, 1000, true},
		testNewParam{1, 1000, true},
		testNewParam{2, 1000, false},
		testNewParam{0.5, 0, false},
		testNewParam{0.5, 100000000, true},
	}
	for _, p := range params {
		if _, err = NewRandSampler(p.ratio, p.width); (err != nil && p.success) || (err == nil && !p.success) {
			t.Errorf("param(ratio, width = %f, %d), success should be %v, but err = %v", p.ratio, p.width, p.success, err)
		}
	}
}

func TestRandSamplerCorrectness(t *testing.T) {
	type Param struct {
		ratio float32
		width uint32
	}
}
