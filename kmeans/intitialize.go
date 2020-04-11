package kmeans

import (
	"math/rand"
)

// InitialAssign randomly chooses first elem of each cluster
func InitialAssign(rawObs [][]int, count int) (map[int][][]int, [][]int){
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
	return clusters, obs
}