package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"grpc_test/proto/grpc/calculator"
)

type calculatorServer struct {
	// Embed the UnimplementedServer to satisfy the interface
	calculator.UnimplementedCalculatorServer
}

func (s *calculatorServer) Add(stream calculator.Calculator_AddServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		num1 := req.Num1
		num2 := req.Num2
		result := num1 + num2
		resp := &calculator.AddResponse{Result: result}

		if err := stream.Send(resp); err != nil {
			return err
		}

		time.Sleep(time.Second) // 模拟耗时操作
	}
}

func (s *calculatorServer) SendInstruction(stream calculator.Calculator_SendInstructionServer) error {
	for {
		// Send an instruction to the client
		resp := &calculator.InstructionResponse{
			Message: "Do something on the client side!",
		}

		if err := stream.Send(resp); err != nil {
			return err
		}

		time.Sleep(time.Second * 3) // 模拟每隔3秒发送一次指令
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	calculator.RegisterCalculatorServer(server, &calculatorServer{})

	fmt.Println("Server is listening on port 50051...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
