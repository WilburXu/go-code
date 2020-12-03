package main

import (
	"fmt"
	"time"
)

func main() {

	tt, _ := time.Parse("2006-01-02", "2018-07-11")
	fmt.Println(tt.Unix())

	//var header http.Header
	//u, _ := url.Parse(`wss://test-gw.nblive.io/ws`)
	////u, _ := url.Parse(`ws://127.0.0.1:9633/ws`)
	//v := u.Query()
	//v.Set("token", "2b537a3ab24ce1883dbac9d7f518ca041cc6ee9264f10873c9b342ea711c4577")
	//v.Set("h_m", "11")
	//v.Set("h_did", "8035E49B-80B8-434A-A1CF-E065B331713B")
	//u.RawQuery = v.Encode()
	//conn, _, e := websocket.DefaultDialer.Dial(u.String(), header)
	//log.Println(u.String())
	//if e != nil {
	//	log.Panicln(e)
	//}
	//log.Println("conn:", conn)
}