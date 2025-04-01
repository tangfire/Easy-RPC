package main

import (
	"Easy-RPC/Starting_development_02/Get_started_quickly_01"
	"context"
	"github.com/smallnest/rpcx/client"
	"log"
)

func main() {

	// #1
	d, _ := client.NewPeer2PeerDiscovery("tcp@localhost:8972", "")
	// #2
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	// #3
	args := &get_start.Args{
		A: 10,
		B: 20,
	}

	// #4
	reply := &get_start.Reply{}

	// #5
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
