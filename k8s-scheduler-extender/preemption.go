package k8s_scheduler_extender

import (
	"k8s.io/api/core/v1"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
)

type Preemption struct {
	Func func(
		pod v1.Pod,
		nodeNameToVictims map[string]*schedulerapi.Victims,
		nodeNameToMetaVictims map[string]*schedulerapi.MetaVictims,
	) map[string]*schedulerapi.MetaVictims
}

func (b Preemption) Handler(
	args schedulerapi.ExtenderPreemptionArgs,
) *schedulerapi.ExtenderPreemptionResult {
	nodeNameToMetaVictims := b.Func(*args.Pod, args.NodeNameToVictims, args.NodeNameToMetaVictims)
	return &schedulerapi.ExtenderPreemptionResult{
		NodeNameToMetaVictims: nodeNameToMetaVictims,
	}
}
