package main

import "testing"

func TestFitIn(t *testing.T) {
	var testEnv1 = Envelope{10,11}
	var testEnv2 = Envelope{12,12}
	want := true
	got := testEnv1.fitIn(testEnv2)
	if got != want {
		got := false
		t.Errorf("Trying if %#v fits in %#v, expected %t but got %t", testEnv2, testEnv1, want, got)
	}
}
