package indexparamcheck

import (
	"strconv"
	"testing"

	"github.com/milvus-io/milvus/pkg/util/metric"

	"github.com/stretchr/testify/assert"
)

func Test_flatChecker_CheckTrain(t *testing.T) {

	p1 := map[string]string{
		DIM:    strconv.Itoa(128),
		Metric: metric.L2,
	}
	p2 := map[string]string{
		DIM:    strconv.Itoa(128),
		Metric: metric.IP,
	}
	p3 := map[string]string{
		DIM:    strconv.Itoa(128),
		Metric: metric.COSINE,
	}
	p4 := map[string]string{
		DIM:    strconv.Itoa(128),
		Metric: metric.HAMMING,
	}
	p5 := map[string]string{
		DIM:    strconv.Itoa(128),
		Metric: metric.JACCARD,
	}

	cases := []struct {
		params   map[string]string
		errIsNil bool
	}{
		{p1, true},
		{p2, true},
		{p3, true},
		{p4, false},
		{p5, false},
	}

	c := newFlatChecker()
	for _, test := range cases {
		err := c.CheckTrain(test.params)
		if test.errIsNil {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}
