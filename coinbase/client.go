package coinbase

import (
	"log"

	ws "golang.org/x/net/websocket"
)

type client struct {
	conn *ws.Conn
}

// NewClient returns a new websocket client.
func NewClient(endpoint string) (CoinbaseWsClient, error) {
	conn, err := ws.Dial(endpoint, "", "http://localhost/")
	if err != nil {
		return nil, err
	}

	log.Printf("websocket connected to: %s", endpoint)

	return &client{
		conn: conn,
	}, nil
}
