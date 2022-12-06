package main

import (
	"context"
	"fmt"
	"flag"
	"log"
	"net"
	"errors"
	"google.golang.org/grpc"
	"kaiko.io/kaiko"
	"kaiko-server/data"
)


var (
	port = flag.Int("port", 8080, "The server port")
)

// server is used to implement kaiko
type server struct {
	kaiko.UnimplementedKaikoServer
}


// Exists implements kaiko
func (s *server) Exists(ctx context.Context, in *kaiko.ExistsRequest) (*kaiko.ExistsResponse, error) {
	exchangeCode := in.GetExchangeCode()
	exchangePairCode := in.GetExchangePairCode()
	log.Printf("Received Exchange code: %v", exchangeCode)
	log.Printf("Received Exchange Pair code: %v", exchangePairCode)
	//check if the pair is supported 
	existCode := data.DataExist(&exchangeCode, &exchangePairCode)
	switch existCode {
    case 0:
        return &kaiko.ExistsResponse{Exists:kaiko.ExistsResponse_UNKNOWN}, nil
    case 1:
        return &kaiko.ExistsResponse{Exists:kaiko.ExistsResponse_YES}, nil
    case 2:
        return &kaiko.ExistsResponse{Exists:kaiko.ExistsResponse_NO},nil
    }
	return &kaiko.ExistsResponse{}, errors.New("Something wrong happened")
}

func main() {
	flag.Parse()
	data.GetData()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	kaiko.RegisterKaikoServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
