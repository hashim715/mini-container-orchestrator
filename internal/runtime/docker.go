package runtime

import (
	"context"
	"fmt"
	"time"

	"github.com/moby/moby/client"
);

func NewDockerRuntime() (Runtime, error) {
	apiClient, err := client.New(client.FromEnv);

	if err != nil {
		fmt.Printf("failed to connect to the docker due to an error: %v\n",err);
		return nil,err;
	};

	docker_runtime := DockerRuntime{apiClient: apiClient};

	return &docker_runtime,nil;
};

func(d *DockerRuntime) ListContainers(ctx context.Context) ([]string,error) {
	containers, err := d.apiClient.ContainerList(ctx, client.ContainerListOptions{});

	if err != nil {
		return []string{},err;
	};

	container_ids := []string{};

	for _, container := range containers.Items {
		container_ids = append(container_ids, container.ID);
	};

	return container_ids,nil;
};

func (d *DockerRuntime) CreateContainer(image string, command string, env string, name string) (string,error) {
	return "",nil;
};

func (d *DockerRuntime) StartContainer(container_id string) error {
	return nil;
};

func (d *DockerRuntime) StopContainer(container_id string, timeout time.Duration) error {
	return nil;
};

func (d *DockerRuntime) RemoveContainer(container_id string) error {
	return nil;
};

func (d *DockerRuntime) StatusOfContainer(container_id string) (string,error) {
	return "",nil;
};

func (d *DockerRuntime) Close() {
	d.apiClient.Close();
};