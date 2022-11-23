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
var replicas []auction.AuctionServiceClient 
var maxBid = 0
var terminate = 0
// var client_obj auction.AuctionServiceClient

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
	replicas = []auction.AuctionServiceClient{client_obj_1,client_obj_2, client_obj_3}


	defer conn_1.Close()
	defer conn_2.Close()
	defer conn_3.Close()

	go func() {
		for {
			if terminate == 1{
				break
			}

			if time.Now().Hour() == hour && time.Now().Minute() == minutes  && time.Now().Second() == 0{
				log.Print("Timer expired, the auction is over.")
				log.Println(replicas)
				for _, client_obj := range replicas {
					//log.Println("loop")
					response, err := client_obj.EndAuction(context.Background(), &auction.Empty{})
					if err != nil {
						log.Fatalf("could not end the auction: %v", err)
					}

					log.Printf("The winner is: %v with bid of value %v", response.HighestBidderId, response.HighestBid)

					terminate = 1
				}
			}
		}
	}()

	for {
		read_input(id)
	}

}

func read_input(id int32) {
	scanner := bufio.NewScanner(os.Stdin)
	log.Println("Type the keyword 'bid' to make a new bid or alternatively type 'result' to retrieve the highest bid")

	for {
		scanner.Scan()
		text := scanner.Text()

		switch {
		case text == "result":
			var resultFinal int32
			m := make(map[int32]int32) // keep track of the number of times a number appear
			for i, client_obj := range replicas {

				result_obj := &auction.ResultRequest{}
				result, err := client_obj.Result(context.Background(), result_obj)
				if err != nil {
					ports = remove(ports, i)
					replicas = remove2(replicas,i)
					log.Printf("could not retrieve the result from server at port 800%v the server has been removed: %v", err, i)
					continue
				}
				//check if value exist 
				_, prs := m[result.HighestBid]
				if prs{
					m[result.HighestBid] = m[result.HighestBid] + 1
				}else{
					m[result.HighestBid] = 1
				}
				// log.Println("The highest bid is: ", result)
			}
			var largest int32 = 0
			for key, element := range m {
				if element > largest{
					resultFinal = key
				}
			}
			log.Println("The highest bid is: ", resultFinal)

		case text == "bid":
			log.Println("Plese type the amount you would like to bid in the command line:")
			scanner.Scan()
			input := scanner.Text()
			amount, err := strconv.Atoi(input)
			if amount < maxBid{
				log.Println("The bid value is lower than the previous bid(s), please try again.")
				
			}else{
				maxBid = amount
				if err != nil {
					log.Fatal("The input you provided is not an integer, please try again.")
				}

				bid := &auction.BidRequest{
					Amount:   int32(amount),
					BidderID: id,
				}

				var ack int
				var acks [3]int
				acks =  [3]int{0,0,0}
				for i, client_obj := range replicas {
					acknowledgement, err := client_obj.Bid(context.Background(), bid)
					if err != nil {
						replicas = remove2(replicas,i)
						 log.Println(i)
						log.Println(err)
						continue
					}
					if acknowledgement.Status == 0{
						acks[0]++
					}else if acknowledgement.Status == 1{
						acks[1]++
					}else{
						acks[2]++
					}
					
					

				}
				var biggest int
				biggest = 0
				for i,_ := range acks{
					if acks[i] > biggest{
						ack = i
					}
				}

				if ack == 0{
					log.Println("FAIL")
				}else if ack == 1{
					log.Println("SUCCESS")
				}else{
					log.Println("EXCEPTION")
				}
				
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
	var result []string
	for i := range slice {
		if(i != s){
			result = append(result,slice[i])
		}
		
	} 	
	return result
}

func remove2(slice []auction.AuctionServiceClient , s int) []auction.AuctionServiceClient  {
	var result []auction.AuctionServiceClient 
	for i, client_obj := range slice {
		if(i != s){
			result = append(result, client_obj)
		}
		
	} 	
	return result
}
