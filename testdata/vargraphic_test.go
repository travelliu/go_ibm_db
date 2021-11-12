// 2021/11/12 Bin Liu <bin.liu@enmotech.com>

package main

import (
	"testing"
)

func TestVarGraphicEmptyString(t *testing.T) {
	if err := VarGraphicEmptyString(); err != nil {
		t.Errorf("Error at VarGraphicEmptyString %s", err)
	}
}

func TestVarGraphicNull(t *testing.T) {
	if err := VarGraphicNull(); err != nil {
		t.Errorf("Error at VarGraphicNull %s", err)
	}
}
