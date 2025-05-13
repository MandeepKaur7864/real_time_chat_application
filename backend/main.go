package main

import (
	"fmt"
	"log"
	"net/http"
	"realtime_chat/pkg/websocket_pkg"

	"github.com/gorilla/websocket"
)

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// We'll need to check the origin of our connection
	// this will allow us to make requests from our React
	// development server to here.
	// For now, we'll do no checking and just allow any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

// define our WebSocket endpoint
func serveWs(pool *websocket_pkg.Pool, w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Host)
	fmt.Println("Webscoekt Endpoint Hit")

	// upgrade this connection to a WebSocket
	// connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Fprint(w, "%+v\n", err)
	}

	client := &websocket_pkg.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client

	client.Read()

	// go websocket_pkg.Writer(ws)

	// websocket_pkg.Reader(ws)

	// // listen indefinitely for new messages coming
	// // through on our WebSocket connection
	// reader(ws)
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)

}

func setupRoutes() {
	// http.HandleFunc("/", home)

	// mape our `/ws` endpoint to the `serveWs` function
	// http.HandleFunc("/ws", serveWs)

	pool := websocket_pkg.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

// func home(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Simple Server")
// }
