package main

import (
	"testing"
)

func TestBooler(t *testing.T) {
	if !Booler() {
		t.Fail()
	}
}
