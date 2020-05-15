package algorithm

import (
	"reflect"
	"testing"
)

func TestUpdate(t *testing.T) {
	t.Run("Empty cluster map", testEmptyMap())
	t.Run("Single obs", testSingleObs())
	t.Run("Correct case", testCorrectUpdate())
}

func testEmptyMap() func(t *testing.T) {
	clusterMap := make(map[int][][]int, 0)
	numFeatures := 0
	return func(t *testing.T) {
		_, err := Update(clusterMap, numFeatures)

		if err == nil {
			t.Error("Expected error to not be nil")
		}
	}
}

func testSingleObs() func(t *testing.T) {
	clusterMap := make(map[int][][]int)
	clusterMap[0] = make([][]int, 1)
	clusterMap[0][0] = []int{ -1, 3, 0 }
	numFeatures := 3

	return func(t *testing.T) {
		cmap, err := Update(clusterMap, numFeatures)

		if err != nil {
			t.Error("Expected error to be nil")
		}

		for i := 0; i < len(cmap[0][0]); i++ {
			if cmap[0][0][i] != clusterMap[0][0][i] {
				t.Error("Not an equivalent map")
			}
		}
	}
}

func testCorrectUpdate() func(t * testing.T) {
	// clusters with very similar obs so no updates
	clusterMap := make(map[int][][]int)
	clusterMap[0] = make([][]int, 2)
	clusterMap[1] = make([][]int, 2)

	clusterMap[0][0] = []int{ 0, 1, 1 }
	clusterMap[0][1] = []int{ 1, 2, 2 }
	clusterMap[1][0] = []int{ 10, 11, 10 }
	clusterMap[1][1] = []int{ 11, 12, 11 }
	numFeatures := 3

	return func(t *testing.T) {
		cmap, err := Update(clusterMap, numFeatures)

		if err != nil {
			t.Error("Expected error to be nil")
		}

		if !reflect.DeepEqual(cmap, clusterMap) {
			t.Error("Incorrectly clustered")
		}
	}
}