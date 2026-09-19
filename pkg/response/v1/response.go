package response

type Response struct {
	Message
	Data   any       `json:"data,omitempty"`
	Errors []Message `json:"errors,omitempty"`
}
type Message struct {
	Code   string         `json:"code,omitempty"`
	Params map[string]any `json:"params,omitempty"`
}
