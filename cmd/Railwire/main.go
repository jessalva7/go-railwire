package main

import (
	"github.com/jessalva/go-railwire/api/plans"
	"github.com/jessalva/go-railwire/pkg/fetching"
	"github.com/jessalva/go-railwire/pkg/repository/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func main() {

	logger := log.New(os.Stdout, "railwire-grpc:", log.LstdFlags)
	planRepository := postgres.NewPostgres()
	fetchingService := fetching.NewService(logger, planRepository)

	gs := grpc.NewServer()

	reflection.Register(gs)
	plans.RegisterPlansService(gs, fetchingService)

	lst, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Fatal("Couldn't connect with tcp")
	}

	_ = gs.Serve(lst)
}
