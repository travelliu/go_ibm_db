// 2022/2/17 Bin Liu <bin.liu@enmotech.com>

package main

import "testing"

func TestClobLength1073741824(t *testing.T) {
	if err := ClobLength1073741824(); err != nil {
		t.Errorf("Error at ClobLength1073741824 %s", err)
	}
}
