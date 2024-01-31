package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CalculateChildrenIdRangeFromInstSpec(t *testing.T) {
	tbl := []struct {
		caseName      string
		parentSpec    int
		childrenSpec  int
		parentInstId  int
		expectedStart int
		expectedEnd   int
	}{
		{"1", 1, 1, 1, 1, 1},
		{"2", 1, 2, 1, 1, 2},
		{"3", 2, 1, 1, 1, 1},
		{"4", 2, 1, 2, 2, 2},
		{"5", 2, 5, 1, 1, 3},
		{"6", 2, 5, 2, 4, 5},
		{"7", 1, 3, 1, 1, 3},
		{"8", 10, 19, 1, 1, 2},
		{"9", 10, 19, 9, 17, 18},
		{"10", 10, 19, 10, 19, 19},
		{"e1", -1, 1, 1, 0, 0},
		{"e2", 1, -1, 1, 0, 0},
		{"e3", 1, 1, 0, 0, 0},
		{"e4", 1, 1, 2, 0, 0},
	}
	for i := range tbl {
		t.Run(tbl[i].caseName, func(t *testing.T) {
			start, end := CalculateChildrenIdRangeFromInstSpec(tbl[i].parentSpec, tbl[i].childrenSpec, tbl[i].parentInstId)
			assert.Equal(t, tbl[i].expectedStart, start, "start not equal")
			assert.Equal(t, tbl[i].expectedEnd, end, "end not equal")
		})
	}
}
