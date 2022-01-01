package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	pb "github.com/SebastianVFugmann/GO-Templates/Active_replication/Service"
	"google.golang.org/grpc"
)

var (
	fe frontend
)

func main() {
	//setupLogger()

	fe = frontend{
		clients: make(map[int32]pb.ServiceClient),
		ctx:     context.Background(),
	}

	setupID()

	dialServer()

	fmt.Println("---------- Welcome to the Service ----------")

	//Basic query tool
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		input := strings.Split(strings.ToLower(command), " ")
		//Command here, length if you need to add an argument
		if len(input) > 2 && input[0] == "bid" {
			auctionId, err := strconv.ParseInt(input[1], 10, 32)
			if err != nil {
				fmt.Printf("Not a valid command: \"%v\" not an integer.", input[1])
				continue
			}
			bid, err := strconv.ParseInt(input[2], 10, 32)
			if err != nil {
				fmt.Printf("Not a valid command: \"%v\" not an integer.", input[2])
				continue
			}
			go fe.bid(int32(auctionId), int32(bid))
		} else if len(input) > 1 && input[0] == "status" {
			auctionId, err := strconv.ParseInt(input[1], 10, 32)
			if err != nil {
				fmt.Printf("Not a valid command: \"%v\" not an integer.", input[1])
				continue
			}
			go fe.status(int32(auctionId))
		} else {
			fmt.Printf("Not a valid command: %v.\n", command)
		}
	}
}

func dialServer() {
	for i := 5000; i <= 5003; i++ {
		address := fmt.Sprintf(":%v", i)
		log.Printf("Dialing %v", address)
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
		if err != nil {
			log.Printf("Could not connect to server: %v\n Error: %v\n", address, err)
			continue
		}
		client := pb.NewAuctionServiceClient(conn)
		fe.clients[int32(i)] = client
	}
	fe.ctx = context.Background()
}

func setupLogger() {
	//Checks of a log file exists and then appends or creates a new
	file, logerr := os.OpenFile("client_logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if logerr != nil {
		log.Fatal(logerr)
	}
	//Sets the loggers output to the log file
	log.SetOutput(file)
}

func setupID() {
	// Asks client for a nickname
	fmt.Println("Please enter your name:")
	reader := bufio.NewReader(os.Stdin)
	tempname, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Failed to read from console: %v \n", err)
	}
	fe.name = strings.Trim(tempname, "\r\n")
	log.Printf("Name saved: %v\n", fe.name)
}
