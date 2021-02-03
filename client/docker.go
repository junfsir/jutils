package client

import (
	"context"
	"os"
	"fmt"

	dockerclient "github.com/docker/docker/client"
)

var Cli *dockerclient.Client

// init docker client
func init()  {
	version := os.Getenv("DOCKER_API_VERSION")

	if len(version) == 0 {
		version = "1.37"
	}
	cli, err := dockerclient.NewClientWithOpts(dockerclient.WithVersion(version))
	if err != nil {
		panic(err)
	}

	Cli = cli
}

// GetCgroupParent gets the parent path of cgroup
func GetCgroupParent(cid string) (string, error) {
	ctx := context.Background()
	inspect, err := Cli.ContainerInspect(ctx, cid)
	if err != nil {
		return "", err
	}
	return inspect.HostConfig.CgroupParent, nil
}

// GetCgroupPath gets the cgroup path of container
func GetCgroupPath(cid, resourceType string) (cgroupPath, parentCgroupPath string, err error) {
	cgroupParent, err := GetCgroupParent(cid)
	if err != nil {
		return
	}
	cgroupPath = fmt.Sprintf("/sys/fs/cgroup/%v%s/%s", resourceType, cgroupParent, cid)
	parentCgroupPath = fmt.Sprintf("/sys/fs/cgroup/%v%s", resourceType, cgroupParent)
	return
}


