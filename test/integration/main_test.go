package integration

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	dockerClient "github.com/docker/docker/client"

	"github.com/aliml92/typesense/typesense"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var (
	baseClient *typesense.Client
	client     *typesense.Client
)

func TestMain(m *testing.M) {
	var err error
	tsDataDir := "typesense-data"
	if err = createDataDir(tsDataDir); err != nil {
		log.Fatal(err)
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %v", err)
	}
	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %v", err)
	}
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "typesense/typesense",
		Tag:        "0.25.1",
		Cmd: []string{
			"--data-dir=/tmp/typesense-data",
			"--api-key=xyz",
			"--enable-cors",
		},
		Mounts: []string{
			"typesense-data:/tmp/typesense-data",
		},
	}, func(hostConfig *docker.HostConfig) {
		hostConfig.AutoRemove = true
		hostConfig.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %v", err)
	}

	time.Sleep(5 * time.Second) // give the server few seconds to get ready

	serverURL := fmt.Sprintf("http://%s", resource.GetHostPort("8108/tcp"))
	baseClient, err = typesense.NewClient(nil, serverURL)
	if err != nil {
		log.Fatalf("Could not create Typesense client: %v", err)
	}
	client = baseClient.WithAPIKey("xyz")

	if err = client.Ping(); err != nil {
		log.Fatalf("Could not connect to Typesense: %v", err)
	}

	code := m.Run()

	// Cleanup; typesense-data directory, docker container and volume
	if err := removeDataDir(tsDataDir); err != nil {
		log.Fatal(err)
	}

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %v", err)
	}

	if err := deleteVolume("typesense-data"); err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}

func createDataDir(dirName string) error {
	if err := removeDataDir(dirName); err != nil {
		return err
	}

	return os.Mkdir(dirName, 0o777)
}

func removeDataDir(dirName string) error {
	if _, err := os.Stat(dirName); !os.IsNotExist(err) {
		return os.RemoveAll(dirName)
	}

	return nil
}

func deleteVolume(volumeName string) error {
	ctx := context.Background()
	cli, err := dockerClient.NewClientWithOpts(
		dockerClient.FromEnv, dockerClient.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	return cli.VolumeRemove(ctx, volumeName, true)
}
