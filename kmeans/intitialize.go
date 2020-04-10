package kmeans

import (
	"math/rand"

	"github.com/ShababKarim/go-cluster/kmeans/util"
)

// Assign randomly chooses first elem of each cluster
func Assign(obs [][]int, count int) [][]int{
	util.ObsPlusCluster(obs)

	for i := 0; i < count; i++ {
		obs[rand.Intn(len(obs))][2] = rand.Intn(count) + 1
	}
	return obs
}