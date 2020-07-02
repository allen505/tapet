package main

import "testing"

func TestSum(t *testing.T){
	total := successSum(1,2)
	if total != 3 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 3)
	}
}

func TestSumWillFail(t *testing.T){
	total := failSum(1,2)
	if total != 3 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 3)
	}
}