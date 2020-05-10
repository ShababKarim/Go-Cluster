package lib

// ClusterMap contains clustered observations in map-like fashion and
// contains some data on cluster validation like WSS, silhouette coeff.
type ClusterMap struct {
	silCoeff int
	wss int
	Obs map[int][][]int
}