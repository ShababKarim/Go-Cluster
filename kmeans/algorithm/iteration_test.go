package algorithm

import (
	"testing"
)

func TestNearestCentroid(t *testing.T) {
	t.Run("empty centroid list", testEmptyList())
	t.Run("empty observation", testEmptyObservation())
	t.Run("perfect match", testPerfectMatch())
}

func testEmptyList() func(t *testing.T) {
	centroidList := make([][]int, 0)
	obs := []int{1, 2, 3}
	numFeatures := 3
	return func(t *testing.T) {
		_, err := NearestCentroid(centroidList, obs, numFeatures)
		if err == nil {
			t.Error("Expected error to not be nil")
		}
	}
}

func testEmptyObservation() func(t *testing.T) {
	centroidList := make([][]int, 2)
	centroidList[0] = []int{ 3, 1 }
	centroidList[1] = []int{ 4, 5 }
	obs := make([]int, 0)
	numFeatures := 2
	return func(t *testing.T) {
		_, err := NearestCentroid(centroidList, obs, numFeatures)
		if err == nil {
			t.Error("Expected error to not be nil")
		}
	}
}

func testPerfectMatch() func(t *testing.T) {
	centroidList := make([][]int, 2)
	centroidList[0] = []int{ 3, 1 }
	centroidList[1] = []int{ 4, 5 }
	obs := []int{ 4, 5 }
	numFeatures := 2
	return func(t *testing.T) {
		index, err := NearestCentroid(centroidList, obs, numFeatures)
		if err != nil {
			t.Error("Expected error be nil")
		}

		if index != 1 {
			t.Error("Returned wrong index")
		}
	}
}