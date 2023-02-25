package main

import "testing"

func Test_main(t *testing.T) {
	var flag bool = false
	if flag {
		t.Skip("スキップします")
	}
	err := readFile("testdata/engine.conf")
	if err != nil {
		t.Errorf(err.Error())
	}
}
