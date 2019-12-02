package main

import (
	"reflect"
	"testing"
)

func TestExecute(t *testing.T) {
	p := &Program{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	expected := &Program{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}
	p.execute()
	if !reflect.DeepEqual(*p, *expected) {
		t.Errorf("wrong final state expected %v, got %v", expected, p)
	}
}
