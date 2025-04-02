package main

import (
	model "Easy-RPC/Starting_development_02/Functions_for_services_02"
	"context"
	"flag"
	"github.com/smallnest/rpcx/server"
)

func mul(ctx context.Context, args *model.Args, reply *model.Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()
	s.RegisterFunction("a.fake.service", mul, "")
	s.Serve("tcp", ":8031")
}
