package lib

import (
	"errors"
)

// NumFeatures returns p as count of features
func NumFeatures(obs [][]int) (int, error) {
	if len(obs) < 0 {
		return -1, errors.New("No observations")
	}
	p := len(obs[0])

	for _, indivObs := range obs {
		if len(indivObs) != p {
			return -1, errors.New("Uneven number of features in the observations")
		}
	}
	
	return p, nil
}