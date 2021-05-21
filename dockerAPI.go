package main

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func listContainers() error {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return (err)
	}

	containers, err := cli.ContainerList(context.Background(),
		types.ContainerListOptions{})
	if err != nil {
		return err
	}
	for {
	}





}
