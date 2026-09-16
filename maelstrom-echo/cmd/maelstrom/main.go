package main

import (
	"encoding/json"
	"log"
	"net/rpc"

	internal "maelstrom-echo/internal/maelstrom-handlers"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	// connect to the local RPC server
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dialing error:", err)
	}
	defer client.Close()

	log.Println("connected to rpc server")
	
	n := maelstrom.NewNode()

	n.Handle("echo", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "echo_ok"

		return n.Reply(msg, body)
	})

	n.Handle("generate", internal.HandleGenerate(n, client))

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
