package application

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"pinger/model"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

const (
	sleepInterval = time.Second * 30
	pingTimeout   = time.Second * 2
	apiUrl        = "http://backend:8382/api/add_status"
)

type Pinger struct {
	status *model.PingStatus
}

func NewApp() *Pinger {
	return &Pinger{
		status: &model.PingStatus{},
	}
}

func (p *Pinger) Run() {
	for {
		if err := p.checkContainerStatus(); err != nil {
			log.Printf("Error: Container check failed: %s", err)
		}
		time.Sleep(sleepInterval)
	}
}

func (p *Pinger) checkContainerStatus() error {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return err
	}

	for _, v := range containers {
		ip := v.NetworkSettings.Networks["container-monitoring_services-network"].IPAddress
		if ip == "" {
			continue
		}
		p.status.IPAddress = ip

		port := v.Labels["ping_port"]
		if port == "" {
			continue
		}

		pingTime, success := pingHost(ip, port)
		p.status.PingTime = pingTime
		if success {
			p.status.LastSuccess = time.Now()
		}

		var jsonData []byte
		if jsonData, err = json.Marshal(p.status); err != nil {
			return err
		}

		resp, err := http.Post(apiUrl, "application/json", bytes.NewReader(jsonData))
		if err != nil {
			log.Printf("Failed to send status: %v", err)
			continue
		}
		resp.Body.Close()
	}

	return nil
}

func pingHost(host, port string) (time.Time, bool) {
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", host, port), pingTimeout)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	return time.Now(), err == nil
}
