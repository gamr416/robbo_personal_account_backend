package docker_client

import (
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/spf13/viper"
	"log"
)

func NewTestDockerClient() (testDockerClient dockertest.Pool) {
	pool, poolErr := dockertest.NewPool("")
	if poolErr != nil {
		log.Fatalf("Could not construct pool: %s", poolErr)
	}

	if err := pool.Client.Ping(); err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	opts := dockertest.RunOptions{
		Name:         viper.GetString("docker.containerName"),
		Repository:   "postgres",
		Tag:          "13",
		Env:          viper.GetStringSlice("docker.environment"),
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "localhost", HostPort: "5433"},
			},
		},
	}

	_, err := pool.RunWithOptions(&opts,
		func(config *docker.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		},
	)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	return
}
