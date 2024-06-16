package ws

import (
	"net/http"
	"time"

	"github.com/lxzan/gws"
)

func Run() {
	upgrader := gws.NewUpgrader(&Handler{}, &gws.ServerOption{
		ParallelEnabled:   true,                                 // Parallel message processing
		Recovery:          gws.Recovery,                         // Exception recovery
		PermessageDeflate: gws.PermessageDeflate{Enabled: true}, // Enable compression
	})
	http.HandleFunc("/connect", func(writer http.ResponseWriter, request *http.Request) {
		socket, err := upgrader.Upgrade(writer, request)
		if err != nil {
			return
		}
		go func() {
			socket.ReadLoop() // Blocking prevents the context from being GC.
		}()
	})
	http.ListenAndServe(":6666", nil)
}

const (
	PingInterval = 5 * time.Second
	PingWait     = 10 * time.Second
)

type Handler struct{}

func (c *Handler) OnOpen(socket *gws.Conn) {
	_ = socket.SetDeadline(time.Now().Add(PingInterval + PingWait))
}

func (c *Handler) OnClose(socket *gws.Conn, err error) {}

func (c *Handler) OnPing(socket *gws.Conn, payload []byte) {
	_ = socket.SetDeadline(time.Now().Add(PingInterval + PingWait))
	_ = socket.WritePong(nil)
}

func (c *Handler) OnPong(socket *gws.Conn, payload []byte) {}

func (c *Handler) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()
	socket.WriteMessage(message.Opcode, message.Bytes())
}
