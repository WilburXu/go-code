package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3/concurrency"
	"log"
	"time"
)

func main() {
	InitEtcd()
	//Mutex()
	for i := 1; i <= 10; i++ {
		go SetMutex(i)
	}

	for {

	}
}

func SetMutex(i int) {
	log.Printf("this is %d start. \n", i)
	var ctxTimeout int64 = 30
	client := GetClient()
	etcdCtx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(ctxTimeout)*time.Second)
	defer cancelFunc()

	etcdSession, mutexObj, err := client.SetMutex(etcdCtx, ctxTimeout, "/xst")
	log.Printf("this is %d get lock. \n", i)

	defer etcdSession.Close()
	defer mutexObj.Unlock(etcdCtx)

	if err != nil {

	}

	log.Println("is sleep")
	time.Sleep(2 * time.Second)
	log.Println("is close")

	log.Printf("this is %d end. \n", i)
}

func Mutex() {
	ctxTimeout := 30
	client := GetClient()
	etcdTimeoutCtx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(ctxTimeout)*time.Second)
	defer cancelFunc()

	response, e := client.Grant(etcdTimeoutCtx, int64(ctxTimeout))
	if e != nil {
		log.Fatal(e.Error())
	}

	etcdSession, err := concurrency.NewSession(client.Client, concurrency.WithLease(response.ID))
	if err != nil {

	}
	defer etcdSession.Close()

	log.Println("is NewMutex")
	mutexObj := concurrency.NewMutex(etcdSession, "/distributed-lock")
	//ctx := context.Background()

	if err := mutexObj.Lock(etcdTimeoutCtx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Do some work in")
	time.Sleep(5 * time.Second)

	if err := mutexObj.Unlock(etcdTimeoutCtx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("released lock for ")
}


//func put() {
//	var err error
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
//	ret, err := client.Put(ctx, "/demo/demo1_key", "demo1_value111")
//	cancel()
//	if err != nil {
//		log.Println(err)
//	}
//
//	log.Printf("put ret: %+v \n", ret)
//}
//
//func get() {
//	var err error
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	resp, err := client.Get(ctx, "/demo/demo1_key")
//	cancel()
//
//	// Get查询还可以增加WithPrefix选项，获取某个目录下的所有子元素
//	// eg: resp, err := client.Get(ctx, "/demo/", clientv3.WithPrefix())
//	if err != nil {
//		fmt.Println("get failed err:", err)
//		return
//	}
//
//	log.Printf("get resp %v \n", resp)
//	log.Printf("get resp kvs %v \n", resp.Kvs)
//
//	// Kvs 返回key的列表
//	for _, item := range resp.Kvs {
//		fmt.Printf("get resp ret %s : %s \n", item.Key, item.Value)
//	}
//}
//
//func delete() {
//	ctx, _ := context.WithTimeout(context.Background(), time.Second)
//	resp, err := client.Delete(ctx, "/demo/demo1_key")
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(resp.PrevKvs)
//}