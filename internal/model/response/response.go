package response

import "hermes/internal/model/response/responsecode"

type WebsocketResponseType string

const (
	WebsocketResponseTypeAck       WebsocketResponseType = "ack"
	WebsocketResponseTypeError     WebsocketResponseType = "error"
	WebsocketResponseTypeGameState WebsocketResponseType = "game_state"
	WebsocketResponseTypeJoin      WebsocketResponseType = "join"
)

type APIResponse struct {
	Code      responsecode.ResponseCode `json:"code"`
	CodeName  string                    `json:"codename"`
	Data      interface{}               `json:"data,omitempty"`       // Optional for success responses
	Error     string                    `json:"error,omitempty"`      // Optional for error responses
	RequestId string                    `json:"request_id,omitempty"` // Optional for now
	Type      WebsocketResponseType     `json:"type,omitempty"`       // for websocket responses
}
