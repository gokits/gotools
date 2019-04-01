package gotools

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type RandSampler struct {
	r     *rand.Rand
	base  int32
	ratio int32
}

func NewRandSampler(ratio float64) (*RandSampler, error) {
	if ratio < 0 || ratio > 1 {
		return nil, errors.New(fmt.Sprintf("ratio %f invalid, must between 0 and 1", ratio))
	}
	s := rand.NewSource(time.Now().UnixNano())
	return &RandSampler{
		r:     rand.New(s),
		base:  10000,
		ratio: int32(ratio * 10000),
	}, nil
}

func (rs *RandSampler) Sample() bool {
	return rs.r.Int63n(int64(rs.base)) < int64(rs.ratio)
}
