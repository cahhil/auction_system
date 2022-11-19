package main

import (
	auction "auction_system/proto"
	"bufio"
	"context"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
)

type Client struct {
	auction.AuctionServiceClient
	ctx context.Context
	id  int32
}

var server auction.AuctionServiceClient

func main() {

	args, _ := strconv.ParseInt(os.Args[1], 10, 32)
	id := int32(args)

	conn, err := grpc.Dial("server", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := auction.NewAuctionServiceClient(conn)
	log.Println("The client is now connected to the server", client)

	client_obj := &Client{
		ctx: context.Background(),
		id:  id,
	}

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		log.Println("Type the keyword 'bid' to make a new bid or alternatively type 'result' to retrieve the highest bid")

		for {
			scanner.Scan()
			text := scanner.Text()

			switch {
			case text == "result":

				result_obj := &auction.ResultRequest{}
				result, err := client_obj.Result(client_obj.ctx, result_obj)
				if err != nil {
					log.Fatalf("could not retrieve the result: %v", err)
				}
				log.Println("The highest bid is: ", result)

			case text == "bid":
				log.Println("Plese type the amount you would liek to bid in the command line:")
				scanner.Scan()
				input := scanner.Text()
				amount, err := strconv.Atoi(input)
				if err != nil {
					log.Fatal("The input you provided is not an integer, pelse try again.")
				}

				bid := &auction.BidRequest{
					Amount:   int32(amount),
					BidderId: client_obj.id,
				}

				acknoledgement, err := server.Bid(context.Background(), bid)

				if err != nil {
					log.Printf("Bid failed:")
					log.Println(err)
				}

				log.Println("Bid response: ", acknoledgement)

			default:
				log.Printf("The input inserted is not correct\n")
			}
		}
	}()
	for {

	}
}
