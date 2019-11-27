package gotools

import (
	"encoding/json"
	"testing"
	"time"
)

type TestDuration struct {
	Timeout Duration
}

func TestDurationUnmarshal(t *testing.T) {
	rawbytes := []byte(`{"Timeout": "10s"}`)
	var td TestDuration
	if err := json.Unmarshal(rawbytes, &td); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if td.Timeout.Duration != 10*time.Second {
		t.Fatalf("unmarshal failed. value of Timeout wrong: %s", td.Timeout)
	}
}
