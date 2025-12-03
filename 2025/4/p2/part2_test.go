package p2

import (
	"testing"
)

func TestSolve(t *testing.T) {
	t.Run("Solution", func(t *testing.T) {
		//Arrange

		//Act

		//Assert
		got := 10
		expected := 3121910778619

		if got != expected {
			t.Errorf("Expexted: %d, Got: %d", expected, got)
		}
	})
}
