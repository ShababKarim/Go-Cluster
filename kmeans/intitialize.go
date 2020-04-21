package kmeans

import (
	"errors"
	"math/rand"
)

// InitialAssign randomly chooses first elem of each cluster
func InitialAssign(rawObs [][]int, count int) (map[int][][]int, [][]int, error){
	var err error

	if len(rawObs) == 0 {
		err = errors.New("No observations")
	} else if len(rawObs) < count || count < 1 {
		err = errors.New("Invalid cluster number")
	}
	
	obs := make([][]int, len(rawObs))
	copy(obs, rawObs)
	clusters := make(map[int][][]int)

	for i := 0; i < count; i++ {
		// select obs
		index := rand.Intn(len(obs))
		// create a list at cluster
		clusterObs := append(make([][]int, 0), obs[index])
		clusters[i] = clusterObs
		// remove from obs slice
		obs = append(obs[0:index], obs[index+1:]...)
	}
	return clusters, obs, err
}