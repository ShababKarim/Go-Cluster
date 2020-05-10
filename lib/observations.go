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
	
	return p, nil
}