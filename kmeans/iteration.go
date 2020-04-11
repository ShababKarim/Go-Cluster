package kmeans

import (
	"errors"
	"math"
)

// CentroidList computes centroids of each cluster
func CentroidList(clusters map[int][][]int, numFeatures int) ([][]int, error){
	var err error
	centroids := make([][]int, len(clusters))
	
	if len(clusters) < 1 {
		err = errors.New("Less than 1 cluster(s)")
	} else if numFeatures < 1 {
		err = errors.New("Not enough features")
	}

	for clusterNum, clusterObs := range clusters {
		averageVector := make([]int, numFeatures)

		for _, vector := range clusterObs {
			for p := 0; p < numFeatures; p++ {
				averageVector[p] += vector[p]
			}
		}

		for p := 0; p < numFeatures; p++ {
			averageVector[p] /= len(clusterObs)
		}

		centroids[clusterNum] = averageVector
	}

	return centroids, err
}

// NearestCentroid finds closest centroid, returns
// corresponding index in clusterMap
func NearestCentroid(centroids [][]int, obs []int, numFeatures int) (int, error){
	closestIndex := -1
	closestDistance := int(^uint(0) >> 1)
	var err error

	for i := 0; i < len(centroids); i++ {
		distance := 0
		for p := 0; p < numFeatures; p++ {
			distance += math.Pow(math.Abs(centroids[i][p] - obs[p]), 2)
		}
		distance = int(math.Sqrt(distance))

		if distance < closestDistance {
			closestDistance = distance
			closestIndex = i
		}
	}
	
	if closestIndex == -1 {
		err = errors.New("No nearest centroid")
	}

	return closestIndex, err
}