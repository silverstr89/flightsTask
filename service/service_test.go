package service

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestFlights(t *testing.T) {
	tests := [][][]string{{{"SFO", "EWR"}}, {{"ATL", "EWR"}, {"SFO", "ATL"}},
		{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}}}

	result := [2]string{"SFO", "EWR"}

	for _, v := range tests {
		r := analyzeFlights(v)
		assert.Equal(t, r, result)
	}
}
