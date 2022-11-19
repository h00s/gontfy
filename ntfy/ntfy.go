package ntfy

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Ntfy struct {
	Server string `json:"server"`
}

func New() *Ntfy {
	return &Ntfy{
		Server: "https://ntfy.sh/",
	}
}

func (n *Ntfy) Send(topic, title, message string) error {
	m, err := json.Marshal(&Message{topic, title, message})
	log.Println(string(m))

	if err != nil {
		return err
	}

	resp, err := http.Post(n.Server, "application/json", bytes.NewBuffer(m))
	log.Println(resp.StatusCode)

	return err
}
