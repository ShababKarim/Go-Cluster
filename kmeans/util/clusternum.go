package util

// ObsPlusCluster add to observations a field for cluster number
// x, y, clusterNumber - initially 0
func ObsPlusCluster(obs [][]int) [][]int{
	for i := 0; i < len(obs); i++ {
		obs[i] = append(obs[i], 0)
	}
	return obs
}