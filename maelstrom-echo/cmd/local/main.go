package main

import _rpc "maelstrom-echo/internal/rpc"

func main() {
	ls := _rpc.LocalService{}
	ls.Start()
}