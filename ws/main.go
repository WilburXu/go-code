package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
)

func main() {
	log.Println("hello ....")
	var header http.Header
	u, _ := url.Parse(`wss://gw.hilivetech.com/ws-game`)
	//u, _ := url.Parse(`ws://127.0.0.1:9633/ws`)
	v := u.Query()
	v.Set("token", "147e91f89fee942513a04fd46de4f223909c02688be8b83d85506d9ba13b41a1")
	v.Set("h_m", "10011")
	v.Set("h_did", "5ff742182b476b60")
	v.Set("h_did", "5ff742182b476b60")
	v.Set("session_id", "bjl-5ff742182b476b60")

	u.RawQuery = v.Encode()
	conn, _, e := websocket.DefaultDialer.Dial(u.String(), header)
	log.Println(u.String())
	if e != nil {
		log.Panicln(e)
	}
	log.Println("conn:", conn)
	log.Println(conn.ReadMessage())
}