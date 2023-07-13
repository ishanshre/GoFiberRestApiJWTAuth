package helpers

type Message struct {
	Message string `json:"message,omitempty"`
	Limit   int    `json:"limit,omitempty"`
	Offset  int    `json:"offset,omitempty"`
	Data    any    `json:"data,omitempty"`
}
