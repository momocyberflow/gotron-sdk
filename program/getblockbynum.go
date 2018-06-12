package main

import (
	"flag"
	"fmt"
	"github.com/tronprotocol/go-client-api/service"
	"log"
	"strings"
)

func main() {
	grpcAddress := flag.String("grpcAddress", "39.106.178.126:50051",
		"gRPC address: <IP:port> example: -grpcAddress localhost:50051")

	num := flag.Int64("number", 0,
		"number: <block number>")

	flag.Parse()

	if (*num < 0) || (strings.EqualFold("", *grpcAddress) && len(*grpcAddress) == 0) {
		log.Fatalln("./get-block-by-num -grpcAddress localhost" +
			":50051 -number 0")
	}

	client := service.NewGrpcClient(*grpcAddress)
	client.Start()
	defer client.Conn.Close()

	block := client.GetBlockByNum(*num)

	fmt.Printf("block: %v\n", block)
}
