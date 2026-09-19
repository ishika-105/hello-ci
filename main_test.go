package main

import "testing"

func TestAdd(t *testing.T){
	result:=add(4,5)

	if result != 9 {
		t.Errorf("Expected 5 but got %d",result)
	}
}