package tdd

import "testing"

func TestVowelCount(t *testing.T) {
	expected := uint(5)
	actual := VowelCount("I love you")
	if actual != expected {
		t.Errorf("actual %d, expected %d", actual, expected)
	}
}
