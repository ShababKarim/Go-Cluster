package kmeans

import (
	"errors"
)

// KMeans creates a cluster map
type KMeans interface {
	New([][]int, int) (map[int][][]int, error)
}

// New creates a map using observations and number of clusters
func New(rawObs [][]int, count int) (map[int][][]int, error) {
	var err error
	numFeatures := -1

	if len(rawObs) == 0 {
		err = errors.New("No observations")
	} else {
		numFeatures = len(rawObs[0])
	}

	clusters, remainObs, err := InitialAssign(rawObs, count)
	centroids, err := CentroidList(clusters, numFeatures)

	return clusters, err
}
