package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

	pb "github.com/SebastianVFugmann/GO-Templates/Active_replication/Service"
	"github.com/SebastianVFugmann/GO-Templates/LamportClock"
	"google.golang.org/grpc"
)

var (
	myPort string
)

type server struct {
	pb.UnimplementedAuctionServer
	lamport LamportClock.LamportClock
	inc     incrementor
}

type incrementor struct {
	lock         sync.Mutex
	currentValue int32
}

func main() {
	fmt.Print("======Server Started====== \n Please write a port between 5000 - 5003. \n")
	setup()

	start()
}

func setup() {
	reader := bufio.NewReader(os.Stdin)

	addressS, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Failed to read from console: %v \n", err)
	}

	//remove the end string and newline chars from the string
	addressS = strings.Replace(addressS, "\r\n", "", -1)

	//parses it as an int
	addressInt, _ := strconv.Atoi(addressS)
	if addressInt >= 5000 && addressInt <= 5003 {
		s := strconv.Itoa(addressInt)

		myPort = ":" + s
	} else {
		log.Printf("The given port is not in this program's range\n")
		fmt.Printf("Server restarting...\n")
		setup()
	}
}

func start() {
	// Creates listener instance
	lis, err := net.Listen("tcp", myPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Creates server instance
	s := grpc.NewServer()

	//Registers service
	cc := createServer()
	pb.RegisterAuctionServiceServer(s, cc)

	// Listen and serve
	log.Printf("T:%d - server listening at %v", cc.lamport.GetTime(), lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func createServer() *server {
	server := &server{
		inc: incrementor{currentValue: 0},
	}
	server.lamport.Initialize()
	return server
}

/*
func (s *server) Increment(ctx context.Context, req *incservice.IncrementRequest) (*incservice.IncrementReply, error) {
	s.inc.lock.Lock()
	defer s.inc.lock.Unlock()

	rep := &incservice.IncrementReply{ValueBefore: s.inc.currentValue}

	if req.Value > s.inc.currentValue {
		rep.Success = true
		s.inc.currentValue = req.Value
	} else {
		rep.Success = false
	}
	return rep, nil
}*/

/*
type Acknowledgement int32

const (
	ACCEPTED      Acknowledgement = 0
	NOTHIGHESTBID Acknowledgement = 1
	AUCTIONENDED  Acknowledgement = 2
)

func (s *server) Bid(ctx context.Context, m *pb.Message) (*pb.Acknowledgement, error) {
	return nil, nil
}*/
