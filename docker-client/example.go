package main

import (
	"fmt"
	"github.com/dotcloud/docker"
	"log"
)
import dockerc "github.com/fsouza/go-dockerclient"

func main() {
	c, err := dockerc.NewClient("unix:///var/run/docker.sock")
	if err != nil {
		log.Print(err)
	}
	opts := dockerc.CreateContainerOptions{}
	config := docker.Config{Image: "benschw/etcd"}

	container, err := c.CreateContainer(opts, &config)
	if err != nil {
		log.Print(err)
	}
	// log.Printf("%+v", container)

	var port docker.Port
	port = "4001/tcp"
	portBinding := []docker.PortBinding{docker.PortBinding{HostIp: "0.0.0.0", HostPort: "11022"}}

	portBindings := make(map[docker.Port][]docker.PortBinding)
	portBindings[port] = portBinding
	hostConfig := docker.HostConfig{PortBindings: portBindings}
	err = c.StartContainer(container.ID, &hostConfig)
	if err != nil {
		log.Print(err)
	}

	err = c.StopContainer(container.ID, 10)
	if err != nil {
		log.Print(err)
	}

}
