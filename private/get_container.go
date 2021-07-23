package private

import (
	"context"
	"sort"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetContainer(id string) types.Container {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	containerIndex := sort.Search(len(containers), func(i int) bool {
		return containers[i].ID == id
	})
	return containers[containerIndex]
}
