package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"

	proto "github.com/tommy-sho/grpc-loadbalncing/app/client/genproto"
	"google.golang.org/grpc"
)

var (
	gatewayPort string
	port        string
	method      string
)

func init() {
	flag.StringVar(&gatewayPort, "g-host", "gateway:50000", "gateway port")
	flag.StringVar(&port, "port", "55555", "gateway port")
	flag.StringVar(&method, "method", "Greeting", "method name")
}

func main() {

	flag.Parse()

	ctx := context.Background()

	gConn, err := grpc.DialContext(ctx, os.Getenv(gatewayPort), grpc.WithInsecure())
	if err != nil {
		panic(fmt.Errorf("failed to connect with backend server error : %v ", err))
	}
	c := proto.NewGreetingServerClient(gConn)
	s := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
L:
	for s.Scan() {
		n := s.Text()
		fmt.Print("> ")
		switch n {
		case "exit":
			break Ler
		default:
			r, err := c.Greeting(ctx, &proto.GreetingRequest{Name: n})
			if err != nil {
				fmt.Println("failed to call Gateway error : ", err)
				break L
			}
			fmt.Println(r.Message)
		}
	}
}
