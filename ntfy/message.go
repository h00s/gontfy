package ntfy

type Message struct {
	Topic   string `json:"topic"`
	Title   string `json:"title"`
	Message string `json:"message"`
}
