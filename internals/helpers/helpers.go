package helpers

type Message struct {
	MessageStatus string `json:"status,omitempty"`
	Message       string `json:"message,omitempty"`
	Limit         int    `json:"limit,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Data          any    `json:"data,omitempty"`
}
