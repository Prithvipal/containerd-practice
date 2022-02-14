package main

import (
	"context"
	"log"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
)

func main() {
	//err := pullRedisImage()
	err := listImages()
	if err != nil {
		log.Fatal(err)
	}
}

func listImages() error {
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return err
	}
	defer client.Close()
	ctx := namespaces.WithNamespace(context.Background(), "example")
	images, err := client.ListImages(ctx)
	if err != nil {
		return err
	}

	for _, image := range images {
		log.Println("Image::", image.Name())
	}
	return nil
}

func pullRedisImage() error {
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return err
	}
	defer client.Close()
	ctx := namespaces.WithNamespace(context.Background(), "example")
	image, err := client.Pull(ctx, "docker.io/library/redis:alpine")

	if err != nil {
		return err
	}
	log.Printf("Successfully pulled %s image\n", image.Name())
	return nil
}
