package main

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"time"
)

//docker run --name consul1 -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600 consul:latest agent -server -bootstrap-expect 2 -ui -bind=0.0.0.0 -client=0.0.0.0

func init() {
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			log.Println("ha-ha")
		})
		http.ListenAndServe(":9527", mux)
	}()


	go func() {
		mux1 := http.NewServeMux()
		mux1.HandleFunc("/check", consulCheck)
		http.ListenAndServe(fmt.Sprintf(":%d", 8080), mux1)
	}()
}

func main() {
	client, err := NewConsulClient("192.168.25.142")
	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	client.registerServer()

	time.Sleep(time.Second * 3)
	a, b, err := client.consulapi.Agent().Service("serverNode_1", nil)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(err)
	select {}
}

type client struct {
	consulapi *consulapi.Client
}

func NewConsulClient(addr string) (*client, error) {
	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	c, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
		return nil, err
	}
	return &client{consulapi: c}, nil
}

func (c *client) registerServer() {
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "serverNode_1"      // 服务节点的名称
	registration.Name = "serverNode"      // 服务名称
	registration.Port = 9527              // 服务端口
	registration.Tags = []string{"v1000"} // tag，可以为空
	registration.Address = localIP()      // 服务 IP
	registration.Check = &consulapi.AgentServiceCheck{ // 健康检查
		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, 8080, "/check"),
		Timeout:                        "3s",
		Interval:                       "5s",  // 健康检查间隔
		DeregisterCriticalServiceAfter: "30s", //check失败后30秒删除本服务，注销时间，相当于过期时间
	}

	err := c.consulapi.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("register server error : ", err)
	}
}

func localIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

var count int

func consulCheck(w http.ResponseWriter, r *http.Request) {
	s := "consulCheck" + fmt.Sprint(count) + "remote:" + r.RemoteAddr + " " + r.URL.String()
	fmt.Println(s)
	fmt.Fprintln(w, s)
	count++
}
