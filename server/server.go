package main

import (
	auction "auction_system/proto"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	auction.UnimplementedAuctionServiceServer
	ctx context.Context
}

func main() {
	timer := time.NewTimer(60 * time.Second)
	//get the first argument form teh command line

	args, _ := strconv.ParseInt(os.Args[1], 10, 32)
	//convert port from string to integer
	ownPort := int32(args)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//create listener on the provided port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", ownPort))
	//error handling
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//create a new grpc server
	grpcServer := grpc.NewServer()
	//register the node as a server
	auction.RegisterAuctionServiceServer(grpcServer, &server{ctx: ctx})
	log.Println("Server started on port", ownPort)

	//start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	<-timer.C
	log.Println("Timer expired")
}

// the implementation of the Bid method taking a context and a BidRequest as input
// and returning a BidResponse as output
func (s *server) Bid(ctx context.Context, in *auction.BidRequest) (*auction.BidResponse, error) {
	return &auction.BidResponse{}, nil
}

// the implementation of the Result method taking a context and a ResultRequest as input
// and returning a ResultResponse as output
func (s *server) Result(ctx context.Context, in *auction.ResultRequest) (*auction.ResultResponse, error) {
	return &auction.ResultResponse{}, nil
}
