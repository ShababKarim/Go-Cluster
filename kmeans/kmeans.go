package kmeans

import (
	"github.com/ShababKarim/go-cluster/kmeans/algorithm"
	l "github.com/ShababKarim/go-cluster/lib"
	v "github.com/ShababKarim/go-cluster/validation"
)

// New is simply a wrapper for the cluster data
func New(obs [][]int, clusterCount int) (*l.ClusterMap, error) {
	cmap := &l.ClusterMap{}

	clusters, err := algorithm.InitialAssign(obs, clusterCount)
	if err != nil {
		return nil, err
	}

	numFeatures, err := l.NumFeatures(obs)
	if err != nil {
		return nil, err
	}

	cl, err := algorithm.Update(clusters, numFeatures)
	if err != nil {
		return nil, err
	}
	cmap.Obs = cl
	
	wss, err := v.TotalWss(cmap.Obs, numFeatures)
	if err != nil {
		return nil, err
	}
	cmap.Wss = wss
	
	return cmap, nil
}