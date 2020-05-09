package kmeans

// Update iterates until clusters stop changing
func Update(clusters map[int][][]int, obs[][]int, numFeatures int) (map[int][][]int, error){
	var err error

	isChanged := true 
	for isChanged {
		// compute centroid of each cluster
		centroids, err := CentroidList(clusters, numFeatures)
		// identify nearest cluster
		// assign observation to nearest cluster
		// a. check if exists already
		// b. append to cluster
	}

	return clusters, err
}