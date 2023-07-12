package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var connectionMap = make(map[int]*connection)

type connection struct {
	conn *websocket.Conn
}

func InitRouter() {
	http.HandleFunc("/ws", handleWs)
	_ = http.ListenAndServe(":8000", nil)
}

func getRoute1(c *gin.Context) {
	data := map[string]string{
		"status": "ok",
	}
	c.JSON(http.StatusOK, data)
}

func handleWs(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		panic(err)
	}

	client := &connection{conn: ws}

	go client.Write()
}

func (c *connection) Write() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	log.Info().
		Msg("Client Connected")
	for {
		select {
		case event, ok := <-eventChan:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}

			message, _ := json.Marshal(event)
			if err := c.write(websocket.TextMessage, message); err != nil {
				log.Err(err).
					Msg("Websocket connection error while writing")
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte("ping")); err != nil {
				log.Err(err).
					Msg("Websocket connection err ping")
				return
			}
		}
	}
}

func (c *connection) write(mt int, payload []byte) error {
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	return c.conn.WriteMessage(mt, payload)
}
