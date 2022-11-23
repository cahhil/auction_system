package main

import (
	auction "auction_system/proto"
	"bufio"
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

var ports = []string{"8000", "8001", "8002"}

var client_obj auction.AuctionServiceClient

func main() {
	arg, _ := strconv.ParseInt(os.Args[1], 10, 32)
	id := int32(arg)

	hour_minutes := os.Args[2]

	//extract the hour from the command line argument
	hour, err := strconv.Atoi(hour_minutes[0:2])
	if err != nil {
		log.Fatalf("Error parsing hour: %v", err)
	}
	//extract the minutes from the command line argument
	minutes, err := strconv.Atoi(hour_minutes[3:5])
	if err != nil {
		log.Fatalf("Error parsing minutes: %v", err)
	}

	conn_1, client_obj_1 := connect_to_server(0)
	conn_2, client_obj_2 := connect_to_server(1)
	conn_3, client_obj_3 := connect_to_server(2)

	defer conn_1.Close()
	defer conn_2.Close()
	defer conn_3.Close()

	go func() {
		for {
			if time.Now().Hour() == hour && time.Now().Minute() == minutes {
				log.Print("Timer expired, the auction is over.")
				for _, client_obj := range []auction.AuctionServiceClient{client_obj_1, client_obj_2, client_obj_3} {
					response, err := client_obj.EndAuction(context.Background(), &auction.Empty{})
					if err != nil {
						log.Fatalf("could not end the auction: %v", err)
					}

					log.Printf("The winner is: %v with bid of value %v", response.HighestBidderId, response.HighestBid)

					break
				}
			}
		}
	}()

	for {
		read_input(client_obj_1, client_obj_2, client_obj_3, id)
	}

}

func read_input(client_obj_1 auction.AuctionServiceClient, client_obj_2 auction.AuctionServiceClient, client_obj_3 auction.AuctionServiceClient, id int32) {
	scanner := bufio.NewScanner(os.Stdin)
	log.Println("Type the keyword 'bid' to make a new bid or alternatively type 'result' to retrieve the highest bid")

	for {
		scanner.Scan()
		text := scanner.Text()

		switch {
		case text == "result":
			for i, client_obj := range []auction.AuctionServiceClient{client_obj_1, client_obj_2, client_obj_3} {

				result_obj := &auction.ResultRequest{}
				result, err := client_obj.Result(context.Background(), result_obj)
				if err != nil {
					ports = remove(ports, i)
					log.Fatalf("could not retrieve the result from server at port 800%v the server has been removed: %v", err, i)
				}
				log.Println("The highest bid is: ", result)
			}

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

			for _, client_obj := range []auction.AuctionServiceClient{client_obj_1, client_obj_2, client_obj_3} {
				acknoledgement, err := client_obj.Bid(context.Background(), bid)

				if err != nil {
					log.Printf("Bid failed:")
					log.Println(err)
				}
				log.Println(acknoledgement)

			}

		default:
			log.Println("The input inserted is not correct")
		}
	}
}

// just an helper method to connect to the knows servers
// returns the connection and the client object
func connect_to_server(index int) (*grpc.ClientConn, auction.AuctionServiceClient) {

	conn, err := grpc.Dial("localhost:"+ports[index], grpc.WithInsecure())
	if err != nil {
		log.Printf("Did not connect to port %v: %v", ports[index], err)
		ports = remove(ports, index)
	}
	return conn, auction.NewAuctionServiceClient(conn)
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
