package consul

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	consul "github.com/hashicorp/consul/api"
)

type Registery struct {
	client *consul.Client
}

func NewRegistery(addr string) (*Registery, error) {
	config := consul.DefaultConfig()
	config.Address = addr
	client, err := consul.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("error on creating new consul client: %w", err)
	}

	return &Registery{client: client}, nil
}

func (r *Registery) Register(ctx context.Context, instanceID, serviceName, hostPort string) error {
	parts := strings.Split(hostPort, ":")
	if len(parts) != 2 {
		return errors.New("invalid host:port format. Eg: localhost:8081")
	}
	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("invalid port address: %w", err)
	}

	host := parts[0]

	err = r.client.Agent().ServiceRegister(&consul.AgentServiceRegistration{
		Address: host,
		Port:    port,
		ID:      instanceID,
		Name:    serviceName,
		Check: &consul.AgentServiceCheck{
			CheckID: instanceID,
			TTL:     "5s",
		},
	})

	if err != nil {
		return fmt.Errorf("error on register service: %w", err)
	}

	return nil
}

func (r *Registery) Deregister(ctx context.Context, instanceID string, _ string) error {
	if err := r.client.Agent().ServiceDeregister(instanceID); err != nil {
		return fmt.Errorf("error on deregister service: %w", err)
	}

	return nil
}

func (r *Registery) HealthCheck(instanceID string, _ string) error {
	if err := r.client.Agent().UpdateTTL(instanceID, "", "pass"); err != nil {
		return fmt.Errorf("error on health check service: %w", err)
	}

	return nil
}

func (r *Registery) Discover(ctx context.Context, serviceName string) ([]string, error) {
	entries, _, err := r.client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, fmt.Errorf("error on retrive entities: %w", err)
	}
	var instances []string
	for _, entry := range entries {
		instances = append(instances, fmt.Sprintf("%s-%d", entry.Node.Address, entry.Service.Port))
	}

	return instances, nil
}
