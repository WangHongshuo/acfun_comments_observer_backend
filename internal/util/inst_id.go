package util

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/asynkron/protoactor-go/actor"
)

const ActorNameSuffixFmt = "-%v"

func GetInstIdFromPid(pid *actor.PID) (int, error) {
	if pid == nil {
		return 0, fmt.Errorf("pid is nil")
	}
	strList := strings.Split(pid.Id, "-")
	if len(strList) == 0 {
		return 0, fmt.Errorf("pid id format error")
	}
	return strconv.Atoi(strList[len(strList)-1])
}

// id start with 1
// example: parentSpec: 2, childrenSpec: 5
// parent_1: [1, 2, 3], parent_2: [4, 5]
func CalculateChildrenIdRangeFromInstSpec(parentSpec, childrenSpec, parentInstId int) (int, int) {
	if parentSpec <= 0 || childrenSpec <= 0 || parentInstId < 1 || parentInstId > parentSpec {
		return 0, 0
	}

	if childrenSpec < parentSpec {
		childrenSpec = parentSpec
	}

	spec := childrenSpec / parentSpec
	rem := childrenSpec % parentSpec

	// front
	if parentInstId <= rem {
		spec += 1
		start := (parentInstId-1)*spec + 1
		end := start + spec - 1
		return start, end
	}

	start := rem*(spec+1) + (parentInstId-rem-1)*spec + 1
	end := start + spec - 1

	return start, end
}
