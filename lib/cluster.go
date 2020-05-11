package lib

// ClusterMap contains clustered observations in map-like fashion and
// contains some data on cluster validation like WSS, silhouette coeff.
type ClusterMap struct {
	SilCoeff int
	Wss float64
	Obs map[int][][]int
}