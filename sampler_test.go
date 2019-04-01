package gotools

import (
	"testing"
)

func TestNewSampler(t *testing.T) {
	var err error
	type testNewParam struct {
		ratio   float64
		success bool
	}
	params := []testNewParam{
		testNewParam{-1, false},
		testNewParam{0, true},
		testNewParam{0.3, true},
		testNewParam{1, true},
		testNewParam{2, false},
		testNewParam{0.5, true},
	}
	for _, p := range params {
		if _, err = NewRandSampler(p.ratio); (err != nil && p.success) || (err == nil && !p.success) {
			t.Errorf("param(ratio= %f), success should be %v, but err = %v", p.ratio, p.success, err)
		}
	}
}

func TestRandSamplerCorrectness(t *testing.T) {
	type Param struct {
		ratio   float64
		loopcnt int64
		shotcnt int64
	}
	params := []Param{
		Param{0, 1000000, 0},
		Param{1, 1000000, 1000000},
		Param{0.2, 1000000, 200000},
		Param{0.3, 1000000, 300000},
		Param{0.7, 1000000, 700000},
		Param{0.3, 1000000, 300000},
	}
	var i int64

	for _, p := range params {
		s, _ := NewRandSampler(p.ratio)
		shotcnt := 0
		for i = 0; i < p.loopcnt; i++ {
			if s.Sample() {
				shotcnt++
			}
		}
		if float64(shotcnt) > float64(p.shotcnt)*1.01 || float64(shotcnt) < float64(p.shotcnt)*0.99 {
			t.Errorf("param = %+v failed, actual shotcnt = %d, sampler = %+v", p, shotcnt, s)
		} else {
			t.Logf("param = %+v success, actual shotcnt = %d", p, shotcnt)
		}
	}
}
