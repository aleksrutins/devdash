package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type ContainerInfo struct {
	Id    string       `json:"containerId"`
	Names []string     `json:"names"`
	Ports []types.Port `json:"ports"`
}

func ListContainers(rw http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	containerInfoList := make([]ContainerInfo, 0)
	for _, container := range containers {
		containerInfo := ContainerInfo{
			Id:    container.ID,
			Names: container.Names,
			Ports: container.Ports,
		}
		containerInfoList = append(containerInfoList, containerInfo)
	}
	jsonData, err := json.Marshal(containerInfoList)
	if err != nil {
		panic(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonData)
}
