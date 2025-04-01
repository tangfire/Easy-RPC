package main

import (
	"Easy-RPC/Starting_development_02/Get_started_quickly_01"
	"github.com/smallnest/rpcx/server"
)

func main() {
	s := server.NewServer()
	s.RegisterName("Arith", new(get_start.Arith), "")
	s.Serve("tcp", ":8972")
}
