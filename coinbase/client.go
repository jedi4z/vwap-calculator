package coinbase

import (
	"context"
	"encoding/json"
	"log"

	ws "golang.org/x/net/websocket"
	"golang.org/x/xerrors"
)

type client struct {
	conn *ws.Conn
}

// NewClient returns a new websocket client.
func NewClient(endpoint string) (CoinbaseClient, error) {
	conn, err := ws.Dial(endpoint, "", "http://localhost/")
	if err != nil {
		return nil, err
	}

	log.Printf("websocket connected to: %s", endpoint)

	return &client{
		conn: conn,
	}, nil
}

func (c *client) Subscribe(ctx context.Context, pairs []string, receiver chan Response) error {
	if len(pairs) == 0 {
		return xerrors.New("pairs array must have a least one pair")
	}

	subscription := Request{
		Type:       RequestTypeSubscribe,
		ProductIDs: pairs,
		Channels: []Channel{
			{Name: ChannelTypeMatches},
		},
	}

	payload, err := json.Marshal(subscription)
	if err != nil {
		return xerrors.Errorf("failed to marshal subscription: %w", err)
	}

	err = ws.Message.Send(c.conn, payload)
	if err != nil {
		return xerrors.Errorf("failed to send subscription: %w", err)
	}

	var subscriptionResponse Response

	err = ws.JSON.Receive(c.conn, &subscriptionResponse)
	if err != nil {
		return xerrors.Errorf("failed to receive subscription response: %w", err)
	}

	if RequestType(subscriptionResponse.Type) == RequestTypeError {
		return xerrors.Errorf("failed to subscribe: %s", subscriptionResponse.Message)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				err := c.conn.Close()
				if err != nil {
					log.Printf("failed closing ws connection: %s", err)
				}
			default:
				var message Response

				err := ws.JSON.Receive(c.conn, &message)
				if err != nil {
					log.Printf("failed receiving message: %s", err)
					break
				}

				receiver <- Response{
					Type:      message.Type,
					Size:      message.Size,
					Price:     message.Price,
					ProductID: message.ProductID,
				}
			}
		}
	}()

	return nil
}
