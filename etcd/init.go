package main

var (
	etcdClient *Client
)

func InitEtcd() {
	var err error
	etcdClient, err = NewClient()
	if err != nil {
		panic(err)
	}
}


func GetClient() *Client {
	return etcdClient
}

