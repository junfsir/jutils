package client

import (
	"fmt"

	core_v1 "k8s.io/api/core/v1"
)

// GetNode get node object by name
func GetNode(name string) (*core_v1.Node, error) {
	item, exists, err := nodeInformer.GetIndexer().GetByKey(name)
	if err != nil {
		return nil, fmt.Errorf("get node: %s failed, error: %s", name, err)
	} else if exists == false {
		return nil, fmt.Errorf("node %s is not exists", name)
	}

	node, ok := item.(*core_v1.Node)
	if !ok {
		return nil, fmt.Errorf("type assertion is wrong, shoud be node")
	}

	return node, nil
}

// GetHostIP get host ip by node name
func GetHostIP(nodeName string) (string, error) {
	node, err := GetNode(nodeName)
	if err != nil {
		return "", err
	}

	var hostIP string
	for _, item := range node.Status.Addresses {
		hostIP = item.Address
		break
	}

	return hostIP, nil
}
