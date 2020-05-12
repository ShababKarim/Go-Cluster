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

func TestCentroidList(t *testing.T) {
	t.Run("less features than in map", testIncorrectCentroid())
	t.Run("correct case", testCorrectCentroid())
}

func testIncorrectCentroid() func(t *testing.T) {
	// not handling this case
	// err should not be nil
	clusterMap := make(map[int][][]int)
	clusterMap[0] = make([][]int, 1)
	clusterMap[1] = make([][]int, 1)

	clusterMap[0][0] = []int{ 1, 2, 3 }
	clusterMap[1][0] = []int{ 2, 4 }
	numFeatures := 2

	return func(t *testing.T) {
		_, err := CentroidList(clusterMap, numFeatures)
		if err != nil {
			t.Error("Expected error to be nil")
		}
	}
}

func testCorrectCentroid() func(t *testing.T) {
	clusterMap := make(map[int][][]int)
	clusterMap[0] = make([][]int, 2)
	clusterMap[1] = make([][]int, 2)

	clusterMap[0][0] = []int{ 1, 4 }
	clusterMap[0][1] = []int{ 3, 6 }

	clusterMap[1][0] = []int{ 1, 4 }
	clusterMap[1][1] = []int{ 3, 6 }
	numFeatures := 2
	return func(t *testing.T) {
		centroids, err := CentroidList(clusterMap, numFeatures)
		if err != nil {
			t.Error("Expected error to be nil")
		}

		for i := 0; i < numFeatures; i++ {
			if centroids[0][i] != centroids[1][i] {
				t.Error("Incorrect centroid value")
			}
		}
	}
}