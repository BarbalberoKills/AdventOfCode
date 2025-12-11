package p2

import (
	"testing"
)

func TestSolve(t *testing.T) {
	t.Run("Solution", func(t *testing.T) {
		//Arrange

		//Act

		//Assert
		got := Solve("../input_test2.txt")
		expected := 2

		if got != expected {
			t.Errorf("Expexted: %d, Got: %d", expected, got)
		}
	})
}
