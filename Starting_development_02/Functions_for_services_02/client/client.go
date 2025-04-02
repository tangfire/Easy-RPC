package main

import (
	model "Easy-RPC/Starting_development_02/Functions_for_services_02"
	"context"
	"github.com/smallnest/rpcx/client"
	"log"
)

func main() {
	// client.go
	d, _ := client.NewPeer2PeerDiscovery("tcp@localhost:8031", "")
	xclient := client.NewXClient("a.fake.service", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &model.Args{
		A: 10,
		B: 20,
	}

	reply := &model.Reply{}
	err := xclient.Call(context.Background(), "mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
