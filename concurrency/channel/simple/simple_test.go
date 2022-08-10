package simple

import (
	"testing"
)

func TestNewChanOwner(t *testing.T) {
	c := NewChanOwner(10)
	resultStream := c()

	expected := 0
	for result := range resultStream {
		if result != expected {
			t.Errorf("expected %v, got %v", expected, result)
		}
		expected++
	}
}
