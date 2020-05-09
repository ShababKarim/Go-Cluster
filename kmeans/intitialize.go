package kmeans

import (
	"errors"
	"math/rand"
)

// InitialAssign randomly chooses first elem of each cluster
func InitialAssign(obs [][]int, count int) (map[int][][]int, error){
	var err error

	if len(obs) == 0 {
		err = errors.New("No observations")
	} else if len(obs) < count || count < 1 {
		err = errors.New("Invalid cluster number")
	}
	
	clusters := make(map[int][][]int)

	for _, vector := range obs {
		// select random cluster number
		k := rand.Intn(count)
		// assign observation to cluster
		if _, exists := clusters[k]; exists {
			clusters[k] = append(clusters[k], vector)
		} else {
			clusters[k] = append(make([][]int, 0), vector)
		}
	}
	return clusters, err
}