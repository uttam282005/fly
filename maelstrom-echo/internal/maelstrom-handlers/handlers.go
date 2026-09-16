package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/rpc"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type Args struct {}
type Reply struct {
	Id int64
}

func HandleGenerate(n *maelstrom.Node, client *rpc.Client) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		if msg.Type() != "generate" {
			return fmt.Errorf("wrong msg type")
		}
		
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "generate_ok"

		args := Args{}
		reply := Reply{}
		err := client.Call("LocalService.GenerateUID", &args, &reply)
		if err != nil {
			log.Println("from me" + err.Error())
			body["id"] = nil
		}
		
		body["id"] = reply.Id 
		
		return n.Reply(msg, body) 
	}
}