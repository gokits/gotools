package gotools

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type RandSampler struct {
	r     *rand.Rand
	width uint32
	ratio uint32
}

func NewRandSampler(ratio float32, width uint32) (*RandSampler, error) {
	if ratio < 0 || ratio > 1 {
		return nil, errors.New(fmt.Sprintf("ratio %f invalid, must between 0 and 1", ratio))
	}
	if width == 0 {
		return nil, errors.New("width should not be zero")
	}
	s := rand.NewSource(time.Now().UnixNano())
	return &RandSampler{
		r:     rand.New(s),
		width: width,
		ratio: uint32(float64(ratio) * float64(width)),
	}, nil
}

func (rs *RandSampler) Sample() bool {
	return rs.r.Int63n(int64(rs.width)) < int64(rs.ratio)
}
