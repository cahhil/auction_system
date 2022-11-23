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

var client_obj auction.AuctionServiceClient

func main() {
	arg, _ := strconv.ParseInt(os.Args[1], 10, 32)
	id := int32(arg)

	conn, client_obj := connect_to_server()

	defer conn.Close()

	for {
		read_input(client_obj, id)
	}

}

func read_input(client_obj auction.AuctionServiceClient, id int32) {
	scanner := bufio.NewScanner(os.Stdin)
	log.Println("Type the keyword 'bid' to make a new bid or alternatively type 'result' to retrieve the highest bid")

	for {
		scanner.Scan()
		text := scanner.Text()

		switch {
		case text == "result":

			result_obj := &auction.ResultRequest{}
			result, err := client_obj.Result(context.Background(), result_obj)
			if err != nil {
				log.Fatalf("could not retrieve the result: %v", err)
			}
			log.Println("The highest bid is: ", result)

		case text == "bid":
			log.Println("Plese type the amount you would like to bid in the command line:")
			scanner.Scan()
			input := scanner.Text()
			amount, err := strconv.Atoi(input)
			if err != nil {
				log.Fatal("The input you provided is not an integer, please try again.")
			}

			bid := &auction.BidRequest{
				Amount:   int32(amount),
				BidderID: id,
			}

			acknoledgement, err := client_obj.Bid(context.Background(), bid)

			if err != nil {
				log.Printf("Bid failed:")
				log.Println(err)
			}

			log.Println("Bid response: ", acknoledgement)

		default:
			log.Println("The input inserted is not correct")
		}
	}
}

// just an helper method to connect to the server
// returns the connection and the client object
func connect_to_server() (*grpc.ClientConn, auction.AuctionServiceClient) {

	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	return conn, auction.NewAuctionServiceClient(conn)
}
