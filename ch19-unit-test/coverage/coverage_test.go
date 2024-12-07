package testCoverage

import "testing"

func TestBazBaz(t *testing.T) {
	expected := 3
	actual := BazBaz(3)
	if actual != expected {
		t.Errorf("actual %d, expected %d", actual, expected)
	}
}
func TestBazBaz2(t *testing.T) {
	expected := 25
	actual := BazBaz(25)
	if actual != expected {
		t.Errorf("actual %d, expected %d", actual, expected)
	}
}
