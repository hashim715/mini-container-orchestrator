package runtime

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/containerd/errdefs"
	"github.com/moby/moby/api/types/container"
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

func(d *DockerRuntime) ListContainers(ctx context.Context) ([]DockerContainer,error) {
	containers, err := d.apiClient.ContainerList(ctx, client.ContainerListOptions{All: true});

	if err != nil {
		return []DockerContainer{},err;
	};

	results := []DockerContainer{};

	for _, container := range containers.Items {
		container_name := "";
		if (len(container.Names) > 0) {
			container_name = strings.TrimPrefix(container.Names[0],"/");
		};

		results = append(results, DockerContainer{ID:container.ID,Name:container_name,Status: container.Status});
	};

	return results,nil;
};

func (d *DockerRuntime) CreateContainer(ctx context.Context, image string, commands []string, env []string, name string) (string,error) {
	reader, err := d.apiClient.ImagePull(ctx, image, client.ImagePullOptions{});

	if err != nil {
		return "",err;
	};

	defer reader.Close();

	if _,err := io.Copy(os.Stdout, reader); err != nil {
		return "",err;
	};	

	resp, err := d.apiClient.ContainerCreate(ctx, client.ContainerCreateOptions{
		Config: &container.Config{
			Cmd: commands,
			Tty: false,
			Env: env,
		},
		Image: image,
		Name: name,
	});

	if err != nil {
		if errdefs.IsConflict(err) {
			fmt.Printf("A container named '%s' already exists!\n", name);

			containers,err := d.ListContainers(ctx);

			if err != nil {
				return "",err;
			};

			for _, container := range containers {
				if container.Name == name {
					return container.ID,nil;
				};
			};

			return "",fmt.Errorf("no such container found with the given name");
		};
		return "", err;
	};

	return resp.ID, nil;
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