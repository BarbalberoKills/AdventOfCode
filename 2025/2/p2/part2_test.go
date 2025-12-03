package p2

import (
	"testing"
)

func TestSolve(t *testing.T) {
	t.Run("Solution", func(t *testing.T) {
		//Arrange

		//Act

		//Assert
		got := Solve("../input_test.txt")
		expected := 4174379265

		if got != expected {
			t.Errorf("Expexted: %d, Got: %d", expected, got)
		}
	})
}
