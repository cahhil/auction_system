# Auction_system

Make sure to have port 8000, 8001 and 8002 free

Then open five terminals and navigate them all to the folder with the code

To start the servers/replicas type the one of the following in each terminal:
 
```go run server/server.go 8000```
```go run server/server.go 8001```
```go run server/server.go 8002```

To start the clients type one of the following in each terminal:

```go run client/client.go 1 13:05```
```go run client/client.go 2 13:05```

The first argument given is the client id

The second argument given is the time the acution will end

As of now the auction will end at 13:05, but this can be changed to your liking following this notation ```<hour:minutes>```

The system supports multiple clients, but can at most use three servers
