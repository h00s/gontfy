package main

import "github.com/h00s/gontfy/ntfy"

func main() {
	n := ntfy.New()
	n.Send("test", "test", "testing it now")
}
