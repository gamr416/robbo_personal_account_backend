package docker_client

import (
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/spf13/viper"
	"log"
)

func NewTestDockerClient() (testDockerClient dockertest.Pool, cleanerContainer func()) {
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
				{HostIP: "0.0.0.0", HostPort: "5433"},
			},
		},
	}

	resource, err := pool.RunWithOptions(&opts,
		func(config *docker.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = docker.NeverRestart()
		},
	)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	resource.Expire(viper.GetUint("container_lifetime"))
	cleanerContainer = func() {
		// purge the container
		err = pool.Purge(resource)
		if err != nil {
			log.Panicf("Could not purge resource: %s", err)
		}
	}
	return
}
