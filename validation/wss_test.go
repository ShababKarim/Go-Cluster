package validation

import "testing"

func TestTotalWss(t *testing.T) {
	t.Run("Empty map case", func(t *testing.T) {
		cmap := make(map[int][][]int)
		numFeatures := 2

		_, err := TotalWss(cmap, numFeatures)
		if err == nil {
			t.Error("Expected error to not be nil")
		}
	})

	t.Run("Correct wss", func(t *testing.T) {
		cmap := make(map[int][][]int)
		cmap[0] = [][]int{ {1, 2}, {2, 4} }
		cmap[1] = [][]int{ {6, 8}, {7, 11} }

		numFeatures := 2
		expected := 7.5

		total, err := TotalWss(cmap, numFeatures)
		if err != nil {
			t.Fatalf("Err: %s", err)
		}

		if total != expected {
			t.Fatalf("Expected %.3f, got %.3f", expected, total)
		}
	})
}