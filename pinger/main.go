package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/docker/docker/client"
)

const interval = time.Second * 30

type PingStatus struct {
	IPAddress   string    `json:"ip_address"`
	PingTime    time.Time `json:"ping_time"`
	Success     bool      `json:"success"`
	LastSuccess time.Time `json:"last_success"`
}

func main() {
	for {
		if err := checkContainerStatus(); err != nil {
			log.Printf("Error: Container check failed: %s", err)
		}
		time.Sleep(interval)
	}
}

func checkContainerStatus() error {
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

		success := pingHost(ip)
		lastSuccess := time.Time{}
		if success {
			lastSuccess = time.Now()
		}

		status := &PingStatus{
			IPAddress:   ip,
			Success:     success,
			LastSuccess: lastSuccess,
		}

		var jsonData []byte
		if jsonData, err = json.Marshal(status); err != nil {
			return err
		}

		resp, err := http.Post("http://backend:8382/api/add_status", "application/json", bytes.NewReader(jsonData))
		if err != nil {
			log.Printf("Failed to send status: %v", err)
			continue
		}
		resp.Body.Close()
	}

	return nil
}

func pingHost(host string) bool {
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%s:8382", host), time.Second*2)
	return err == nil
}
