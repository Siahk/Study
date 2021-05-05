package main

import (
	"testing"
)

func TestPrint(t *testing.T){
	res := Print1to20()
	if res != 210 {
		t.Errorf("Return value not valid")
	}
}
