package qrys

import (
	"testing"
)

func TestTrim(t *testing.T) {
	a := trim("/user", "/")
	if a != "user" {
		t.Errorf("trim() failed")
	}
}

func TestSplitAndTrim(t *testing.T) {
	a := split(trim("/user/:id", "/"), "/")
	if a[0] != "user" {
		t.Errorf("failed")
	}
	if a[1] != ":id" {
		t.Errorf("failed")
	}
}
