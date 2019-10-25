package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type message struct {
	ID   int
	data []byte
}

var db *sqlx.DB

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func init() {
	var err error
	db, err = sqlx.Connect("sqlite3", "gelibert.db")
	// db, err = sqlx.Connect("mysql", "arthur:Nfnmzyf@tcp(217.12.127.253:3306)/gelibert")
	if err != nil {
		log.Println(err)
	}
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(p), messageType)
		// if err := conn.WriteMessage(messageType, append(p, "Echo"...)); err != nil {
		// 	log.Println(err)
		// 	return
		// }
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected")

	var ordersList Orders
	err = db.Select(&ordersList, "SELECT * FROM orders WHERE courier_id > ?", 15)
	if err != nil {
		log.Println(err)
	}

	err = ws.WriteJSON(&ordersList)
	if err != nil {
		log.Println(err)
	}

	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Start ...")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8787", nil))
}
