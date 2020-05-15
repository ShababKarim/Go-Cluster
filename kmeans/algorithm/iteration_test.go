package algorithm

import (
	"testing"
)

func TestNearestCentroid(t *testing.T) {

	cases := []struct{ 
		name string 
		centroidList [][]int
		obs []int
		numFeatures int
		expected int 
		}{
		{ 
			name: "empty centroid list", 
			centroidList: nil, 
			obs: []int{1,2,3}, 
			numFeatures: 3, 
			expected: -1,
		},
		{ 
			name: "empty observation",
			centroidList: [][]int{{3,1}, {4,5}}, 
			obs: nil,
			numFeatures: 2,
			expected: -1,
		},
		{ 
			name: "perfect match",
			centroidList: [][]int{{3,1}, {4,5}}, 
			obs: []int{4,5}, 
			numFeatures: 2, 
			expected: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := NearestCentroid(
				tc.centroidList,
				tc.obs,
				tc.numFeatures)
			
			if tc.expected == -1 && err == nil{
				t.Error("Expected error to not be nil")
			} else if tc.expected != -1 {
				if err != nil {
					t.Fatalf("Err: %s", err)
				} else if tc.expected != got {
					t.Fatalf("Expected %d, got %d", tc.expected, got)
				}
			}
		})
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