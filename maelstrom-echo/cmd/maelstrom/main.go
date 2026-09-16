package main

import (
	"encoding/json"
	"fmt"
	"log"
	"uuid"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	// connect to the local RPC server
	// client, err := rpc.Dial("tcp", "localhost:1234")
	// if err != nil {
	// 	log.Fatal("Dialing error:", err)
	// }
	// defer client.Close()

	// log.Println("connected to rpc server")

	n := maelstrom.NewNode()

	n.Handle("echo", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "echo_ok"

		return n.Reply(msg, body)
	})

	n.Handle("generate", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "generate_ok"
		body["id"] = fmt.Sprintf("%s-%s", n.ID(), uuid.New().String())

		return n.Reply(msg, body)
	})

	// n.Handle("generate", internal.HandleGenerate(n, client))

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
