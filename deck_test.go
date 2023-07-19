package main

import "testing"

//Always try to name the test functions "Test" and then the actual name of the function you want to test

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}
