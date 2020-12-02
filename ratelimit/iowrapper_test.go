package ratelimit

import (
	"bytes"
	"io"
	"math/rand"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestReader(t *testing.T) {
	b := make([]byte, 10*1024*1024)
	_, _ = rand.Read(b)
	r := TokenBucketReader(bytes.NewReader(b), rate.NewLimiter(1024*1024, 64*1024))
	p := make([]byte, len(b))
	w := bytes.NewBuffer(p)
	now := time.Now()
	tmpbuf := make([]byte, 4*1024)
	n, err := io.CopyBuffer(w, r, tmpbuf)
	elapsed := time.Since(now)
	if err != nil && err != io.EOF {
		t.Fatalf("copy error: %v", err)
	}
	if n != 10*1024*1024 {
		t.Fatalf("n should be %d, actual %d", 10*1024*1024, n)
	}
	if elapsed < 9*time.Second || elapsed > 11*time.Second {
		t.Fatalf("elapsed should be around 10s, actual %dms", elapsed/time.Millisecond)
	}
}

func TestWriter(t *testing.T) {
	out := make([]byte, 10*1024*1024)
	_, _ = rand.Read(out)
	in := make([]byte, len(out))
	r := bytes.NewBuffer(out)
	w := TokenBucketWriter(bytes.NewBuffer(in), rate.NewLimiter(1024*1024, 64*1024))
	now := time.Now()
	tmpbuf := make([]byte, 4*1024)
	n, err := io.CopyBuffer(w, r, tmpbuf)
	elapsed := time.Since(now)
	if err != nil && err != io.EOF {
		t.Fatalf("copy error: %v", err)
	}
	if n != 10*1024*1024 {
		t.Fatalf("n should be %d, actual %d", 10*1024*1024, n)
	}
	if elapsed < 9*time.Second || elapsed > 11*time.Second {
		t.Fatalf("elapsed should be around 10s, actual %dms", elapsed/time.Millisecond)
	}
}
