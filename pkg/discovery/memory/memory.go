package memory

import (
	"context"
	"errors"
	"snapp-food/pkg/discovery"
	"sync"
	"time"
)

type serviceNameType string
type instanceIDType string

type Registery struct {
	sync.RWMutex
	serviceAddrs map[serviceNameType]map[instanceIDType]*serviceInstance
}

type serviceInstance struct {
	hostPort   string
	lastActive time.Time
}

func NewRegistery() *Registery {
	return &Registery{
		serviceAddrs: make(map[serviceNameType]map[instanceIDType]*serviceInstance),
	}
}

func (r *Registery) Register(ctx context.Context, instanceID, serviceName, hostPort string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.serviceAddrs[serviceNameType(serviceName)]; !ok {
		r.serviceAddrs[serviceNameType(serviceName)] = map[instanceIDType]*serviceInstance{}
	}
	r.serviceAddrs[serviceNameType(serviceName)][instanceIDType(instanceID)] = &serviceInstance{
		hostPort:   hostPort,
		lastActive: time.Now(),
	}

	return nil
}

func (r *Registery) Deregister(ctx context.Context, instanceID, serviceName string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.serviceAddrs[serviceNameType(serviceName)]; !ok {
		return discovery.ErrNotFound
	}

	delete(r.serviceAddrs[serviceNameType(serviceName)], instanceIDType(instanceID))
	return nil
}

func (r *Registery) HealthCheck(instanceID, serviceName string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.serviceAddrs[serviceNameType(serviceName)]; !ok {
		return errors.New("service not registered yet")
	}

	if _, ok := r.serviceAddrs[serviceNameType(serviceName)][instanceIDType(instanceID)]; !ok {
		return errors.New("service instance not registered yet")
	}

	r.serviceAddrs[serviceNameType(serviceName)][instanceIDType(instanceID)].lastActive = time.Now()
	return nil
}

func (r *Registery) Discover(ctx context.Context, serviceName string) ([]string, error) {
	r.Lock()
	defer r.Unlock()

	if len(r.serviceAddrs[serviceNameType(serviceName)]) == 0 {
		return nil, discovery.ErrNotFound
	}

	var res []string
	for _, v := range r.serviceAddrs[serviceNameType(serviceName)] {
		if time.Since(v.lastActive) > 5*time.Second {
			continue
		}
		res = append(res, v.hostPort)
	}

	return res, nil
}
