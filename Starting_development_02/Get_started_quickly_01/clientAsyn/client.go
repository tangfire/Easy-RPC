package main

import (
	get_start "Easy-RPC/Starting_development_02/Get_started_quickly_01"
	"context"
	"github.com/smallnest/rpcx/client"
	"log"
)

func main() {
	d, _ := client.NewPeer2PeerDiscovery("tcp@localhost:8972", "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &get_start.Args{
		A: 10,
		B: 20,
	}

	reply := &get_start.Reply{}
	call, err := xclient.Go(context.Background(), "Mul", args, reply, nil)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	replyCall := <-call.Done
	if replyCall.Error != nil {
		log.Fatalf("failed to call: %v", replyCall.Error)
	} else {
		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
	}
}
