// 2022/11/26 Bin Liu <bin.liu@enmotech.com>

package main

import (
	"testing"
)

func TestIntervalString(t *testing.T) {
	if err := IntervalTest(); err != nil {
		t.Errorf("Error at TestIntervalString %s", err)
	}
}