package runtime

import (
	"context"
	"time"

	"github.com/moby/moby/client"
);

type DockerRuntime struct {
	apiClient *client.Client
};

type Runtime interface {
	CreateContainer(image string, command string, env string, name string) (string,error);
	StartContainer(container_id string) error;
	StopContainer(container_id string, timeout time.Duration) error;
	RemoveContainer(container_id string) error;
	StatusOfContainer(container_id string) (string,error);
	ListContainers(ctx context.Context) ([]string,error);
	Close();
};

