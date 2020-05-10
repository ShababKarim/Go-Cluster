package algorithm

// Update iterates until clusters stop changing
func Update(clusters map[int][][]int, numFeatures int) (map[int][][]int, error){
	var e error
	isChanged := true 

	for isChanged {
		isChanged = false
		newClusters := make(map[int][][]int)
		// compute centroid of each cluster
		centroids, err := CentroidList(clusters, numFeatures)
		if err != nil {
			return nil, err
		}
		// each obs - identify nearest cluster
		for clusterNum, clusterObs := range clusters {
			for _, obs := range clusterObs {

				// assign observation to nearest cluster
				nearestClusterNum, err := NearestCentroid(centroids, obs, numFeatures)
				if err != nil {
					return nil, err
				}

				// if diff assignment than before, flag as true
				if nearestClusterNum != clusterNum {
					isChanged = true
				}

				// assign to new cluster
				if _, exists := newClusters[nearestClusterNum]; exists {
					newClusters[nearestClusterNum] = append(newClusters[nearestClusterNum], obs)
				} else {
					newClusters[nearestClusterNum] = append(make([][]int, 0), obs)
				}

				// rewrite old cluster with new
				clusters = newClusters
			}
		}
	}

	return clusters, e
}