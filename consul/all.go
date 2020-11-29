package main

import (
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"log"
	"net/http"
)


type client struct {
	consul *consul.Client
}

//NewConsul returns a Client interface for given consul address
func NewConsulClient(addr string) (*client, error) {
	config := consul.DefaultConfig()
	config.Address = addr
	c, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &client{consul: c}, nil
}

// Register a service with consul local agent
func (c *client) Register(name string, port int) error {
	reg := &consul.AgentServiceRegistration{
		ID:   name,
		Name: name,
		Port: port,
	}
	return c.consul.Agent().ServiceRegister(reg)
}

// DeRegister a service with consul local agent
func (c *client) DeRegister(id string) error {
	return c.consul.Agent().ServiceDeregister(id)
}

// Service return a service
func (c *client) Service(service, tag string) ([]*consul.ServiceEntry, *consul.QueryMeta, error) {
	passingOnly := true
	addrs, meta, err := c.consul.Health().Service(service, tag, passingOnly, nil)
	if len(addrs) == 0 && err == nil {
		return nil, nil, fmt.Errorf("service ( %s ) was not found", service)
	}
	if err != nil {
		return nil, nil, err
	}
	return addrs, meta, nil
}

func init() {
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				log.Println("request:", r)
			})
}

const (
	name = `test-server`
)

func main() {
	cli, err := NewConsulClient("127.0.0.1:8500")
	if err != nil {
		log.Panicln(err)
	}

	entries, meta, err := cli.Service("live-session", "")
	fmt.Println(entries)
	fmt.Println(meta)
	//cli, err := NewConsulClient("127.0.0.1:8500")
	//if err != nil {
	//	log.Panicln(err)
	//}
	//
	//// http server --> 8000 listen
	//go func() {
	//	err := http.ListenAndServe(":8000", nil)
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//}()
	//
	//time.Sleep(time.Second)
	//log.Println("register now...")
	//err = cli.Register(name, 8000)
	//if err != nil {
	//	log.Panicln(err)
	//}
	//log.Println("register ok")
	//
	//// find services
	//time.Sleep(time.Second * 3)
	//log.Println("finding service now...")
	//entries, meta, err := cli.Service(name, "")
	//log.Println(entries, meta, err)
	//s1 := entries[0]
	//log.Printf("entry 0: %v", s1)
	//log.Printf("s1: %s", s1.Service.Address)
	//select {}
}
