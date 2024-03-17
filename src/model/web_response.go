package model

type WebResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}
