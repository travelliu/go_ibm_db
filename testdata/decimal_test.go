// Copyright © 2021 Bin Liu <bin.liu@enmotech.com>

package main

import (
	"testing"
)

func TestDecimalArray(t *testing.T) {
	if DecimalArray() != nil {
		t.Error("Error at DecimalArray")
	}
}
