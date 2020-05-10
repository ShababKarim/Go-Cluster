package kmeans

import (
	"github.com/ShababKarim/go-cluster/kmeans/algorithm"
)

// ClusterMap contains clustered observations in map-like fashion and
// contains some data on cluster validation like WSS, silhouette coeff.
type ClusterMap struct {
	silCoeff int
	wss int
	Obs map[int][][]int
}

// New is simply a wrapper for the cluster data
func New(obs [][]int, clusterCount int) (*ClusterMap, error) {
	cmap := &ClusterMap{}

	clusters, err := algorithm.InitialAssign(obs, clusterCount)
	if err != nil {
		return nil, err
	}

	numFeatures, err := algorithm.NumFeatures(obs)
	if err != nil {
		return nil, err
	}

	cl, err := algorithm.Update(clusters, numFeatures)
	if err != nil {
		return nil, err
	}
	cmap.Obs = cl

	return cmap, nil
}