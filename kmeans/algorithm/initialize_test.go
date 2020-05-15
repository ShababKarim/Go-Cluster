package algorithm

import (
	"testing"
)

func TestInitialAssign(t *testing.T) {
	t.Run("0 condition", testZeroCondition())
	t.Run("less observations than clusters", testInvalidCount())
	t.Run("correctly assigned", testCorrect())	
}

func testZeroCondition() func(t *testing.T) {
	return func(t *testing.T) {
		obs := make([][]int, 0)
		cmap, err := InitialAssign(obs, 5)

		if err == nil && cmap != nil {
			t.Error("Expected err to not be nil")
		}
	}
}

func testInvalidCount() func(t *testing.T) {
	return func(t *testing.T) {
		obs := make([][]int, 2)

		obs[0] = []int{1, 2, 3}
		obs[1] = []int{4, 5, 6} 

		cmap, err := InitialAssign(obs, 5)

		if err == nil && cmap != nil {
			t.Error("Expected err to not be nil")
		}
	}
}

func testCorrect() func(t *testing.T) {
	return func (t *testing.T) {
		obs := make([][]int, 2)

		obs[0] = []int{1, 2, 3}
		obs[1] = []int{4, 5, 6} 

		cmap, err := InitialAssign(obs, 1)

		if err != nil {
			t.Fatalf("Err: %s", err)
		} else if cmap == nil {
			t.Error("Expected cmap to not be nil")
		}

		if len(cmap[0]) != len(obs) {
			t.Error("Not all observations assigned")
		}
	}
}