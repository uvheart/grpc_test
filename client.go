package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"grpc_test/proto/grpc/calculator"
)

func main() {
	conn, err := grpc.Dial("120.92.146.58:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := calculator.NewCalculatorClient(conn)

	// Start a stream for receiving instructions
	instructionStream, err := client.SendInstruction(context.Background())
	if err != nil {
		log.Fatalf("Error creating instruction stream: %v", err)
	}

	// Receive and process instructions from the server
	for {
		resp, err := instructionStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving response: %v", err)
		}

		fmt.Printf("Server instruction: %s\n", resp.Message)

		// TODO: Add your client-side logic to handle the received instruction
	}

	fmt.Println("Instruction stream closed.")
}
