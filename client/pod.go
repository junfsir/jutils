package client

import (
	"fmt"
	core_v1 "k8s.io/api/core/v1"
)

type Pod struct {
	Pod *core_v1.Pod
}

// NewPod get a pod struct
func NewPod(ns, name string) (*Pod, error) {
	key := fmt.Sprintf("%s/%s", ns, name)
	item, exists, err := podInformer.GetIndexer().GetByKey(key)
	if err != nil {
		return nil, fmt.Errorf("get pod: %s failed, error: %s", name, err)
	} else if exists == false {
		return nil, fmt.Errorf("pod: %s is not exists", name)
	}

	pod, ok := item.(*core_v1.Pod)
	if !ok {
		return nil, fmt.Errorf("type assertion is wrong, shoud be pod")
	}

	return &Pod{Pod: pod}, nil
}

// GetContainerResources get container original resources
func (p *Pod) GetContainerResources(containerName string) map[string]int64 {
	var originalLimit = make(map[string]int64)
	for _, container := range p.Pod.Spec.Containers {
		if container.Name == containerName {
			memLimits := container.Resources.Limits["memory"]
			cpuLimits := container.Resources.Limits["cpu"]
			originalLimit["memory"] = memLimits.Value()
			originalLimit["cpu"] = cpuLimits.MilliValue()
			break
		}
	}
	return originalLimit
}

