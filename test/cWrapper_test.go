package test

import "testing"

func TestCallTest(t *testing.T) {
	got := callTest()
	want := 69

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
