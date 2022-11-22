package main

import (
	auction "auction_system/proto"
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	status "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
)

type server struct {
	auction.UnimplementedAuctionServiceServer
	ctx           context.Context
	mutex         sync.Mutex
	port          int32
	highestBid    int32
	highestBidder int32
	clients       []int32
	clientCounter int32
}

func main() {
	timer := time.NewTimer(60 * time.Second)
	//get the first argument form teh command line

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//convert port from string to integer
	ownPort := int32(8000)

	server := &server{
		ctx:        ctx,
		port:       ownPort,
		highestBid: 0,
		mutex:      sync.Mutex{},
		clients:    make([]int32, 5),
	}

	//start the server
	initServer(server)

	<-timer.C
	log.Printf("Timer expired, the auction is over. The winner is client %v with a bid of %v",
		server.highestBidder, server.highestBid)
}

func initServer(server *server) {
	//create listener on the provided port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", server.port))
	//error handling
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//create a new grpc server
	grpcServer := grpc.NewServer()
	//register the node as a server
	auction.RegisterAuctionServiceServer(grpcServer, server)
	log.Println("Server started on port", server.port)

	//start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

// the implementation of the Bid method taking a context and a BidRequest as input
// and returning a BidResponse as output
func (s *server) Bid(ctx context.Context, in *auction.BidRequest) (*auction.BidResponse, error) {

	var status_obj status.Status

	id := in.BidderID

	//check if the client is already registered
	//when creating a client in the client.go
	//set his id to be zero so
	//initizalizatino routine can be run
	if in.BidderID == 0 {
		//update the client id
		//and add it to the list of clients
		s.mutex.Lock()
		s.clientCounter++
		id = s.clientCounter
		s.clients = append(s.clients, id)
		s.mutex.Unlock()

	}

	if in.Amount > s.highestBid {
		//status code 1 means OK
		status_obj = status.Status{
			Code:    1,
			Message: "You're winning, your bid is higher than the highest bid",
		}
		s.mutex.Lock()
		//if the incoming bid is higher than the current bid,
		//update the current bid and the highest bidder
		s.highestBid = in.Amount
		s.clients = append(s.clients, id)
		s.highestBidder = id

		s.mutex.Unlock()
	} else {
		//if the incoming bid is lower than the current bid,
		//do not update any variable and
		//return the current bid and the highest bidder
		//status code 1 means error
		status_obj = status.Status{
			Code:    1,
			Message: "Your bid is lower than the current bid, place a higher one...",
		}
	}

	return &auction.BidResponse{Status: auction.Status(status_obj.Code)}, nil
}

// the implementation of the Result method taking a context and a ResultRequest as input
// and returning a ResultResponse as output
func (s *server) Result(ctx context.Context, in *auction.ResultRequest) (*auction.ResultResponse, error) {

	return &auction.ResultResponse{HighestBidderId: s.highestBidder, HighestBid: s.highestBid}, nil
}
