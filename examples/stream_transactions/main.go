package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	indexerv1 "github.com/justlstn/aptos-grpc-go/aptos/indexer/v1"
)

type authTokenCredential struct {
	token string
}

func (a authTokenCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + a.token,
	}, nil
}

func (a authTokenCredential) RequireTransportSecurity() bool {
	return true
}

func main() {
	apiKey := os.Getenv("APTOS_API_KEY")
	if apiKey == "" {
		log.Fatal("APTOS_API_KEY environment variable not set")
	}

	tlsCreds := credentials.NewTLS(&tls.Config{})
	authCreds := authTokenCredential{token: apiKey}

	conn, err := grpc.Dial(
		"grpc.mainnet.aptoslabs.com:443",
		grpc.WithTransportCredentials(tlsCreds),
		grpc.WithPerRPCCredentials(authCreds),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := indexerv1.NewRawDataClient(conn)

	startVersion := uint64(0)
	req := &indexerv1.GetTransactionsRequest{
		StartingVersion: &startVersion,
	}

	stream, err := client.GetTransactions(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to get transactions: %v", err)
	}

	count := 0
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving transaction: %v", err)
		}

		fmt.Printf("Received %d transactions\n", len(resp.Transactions))
		for _, tx := range resp.Transactions {
			fmt.Printf("  Version: %d, Type: %v\n", tx.Version, tx.Type)
			count++
			if count >= 10 {
				return
			}
		}
	}
}
