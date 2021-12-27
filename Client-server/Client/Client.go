package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/SebastianVFugmann/GO-Templates/Client-server/Proto"
	"github.com/SebastianVFugmann/GO-Templates/LamportClock"
	"github.com/thecodeteam/goodbye"
	"google.golang.org/grpc"
)

const (
	address = "localhost:6969"
)

var (
	myClock LamportClock.LamportClock
	MyId    int32
)

func main() {
	//Checks of a log file exists and then appends or creates a new
	file, logerr := os.OpenFile("client_logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if logerr != nil {
		log.Fatal(logerr)
	}
	//Sets the loggers output to the log file
	log.SetOutput(file)

	myClock.Initialize()
	MyId = 0

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	//Gets the context
	ctx := context.Background()
	//Creates a client through the connection
	c := pb.NewServiceClient(conn)

	//What is to be eecuted when a client is closed
	defer goodbye.Exit(ctx, -1)
	goodbye.Notify(ctx)
	goodbye.RegisterWithPriority(func(ctx context.Context, sig os.Signal) {
		leave(ctx, c)
		conn.Close()
	}, 0)

	createID(ctx, c)

	//Gets the client to join the system and listens for messages, hence it is a go routine
	go join(ctx, c)

	//Waits for the user to enter messages and sends them
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		go method(ctx, c, scanner.Text(), false)
	}
}

//Private method that calls the SendMessage method defined in the proto file
func method(ctx context.Context, client pb.ServiceClient, message string, special bool) {
	//Increment clock since local event happened
	myClock.Increment()

	//Gets the stream returned by the SendMessage method
	stream, err := client.Method(ctx)
	if err != nil {
		log.Printf("Cannot send message: error: %v \n", err)
	}

	//Creates the new message to be sent and sends it
	msg := pb.Object{}
	stream.Send(&msg)

	//Gets the server response returned by the server from the stream
	ack, err2 := stream.CloseAndRecv()
	if err != nil {
		log.Printf("T:%d - %v returned when sending message - %v\n", ack.Time, ack.HttpStatusCode, err2)
		fmt.Printf("T:%d - %v returned when sending message - %v\n", ack.Time, ack.HttpStatusCode, err2)
	}
	myClock.SyncTime(ack.Time)
	log.Printf("T:%d - Message sent: %v \n", ack.Time, ack.HttpStatusCode)
}

//A private method created to call the Join method defined in the proto file
func join(ctx context.Context, client pb.ChittyChatServiceClient) {
	//Increase clock, since new event happened
	myClock.Increment()

	//Creates the clients join request
	join := &pb.JoinRequest{
		Id:   MyId,
		Time: myClock.GetTime(),
	}

	//Gets the stream that is returned by the JoinChannel method defined in the proto
	stream, err := client.Join(ctx, join)
	if err != nil {
		log.Fatalf("client.JoinChannel(ctx, &channel) throws: %v", err)
	}
	//Sleeps to allow the new user to have it's channel created
	fmt.Printf("Connecting to the Chitty Chat...\n")
	time.Sleep(2 * time.Second)

	//Creates an infinite loop that makes sure the stream is always open and waiting for messages
	waitc := make(chan struct{})

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive message from channel joining. \nErr: %v", err)
			}

			if in.Special {
				//Syncs time with the received message
				myClock.SyncTime(in.Time)
				//Prints message in specific format
				fmt.Printf("T:%d - %v %v \n", myClock.GetTime(), in.Id, in.Message)
			} else if MyId != in.Id {
				//Syncs time with the received message
				myClock.SyncTime(in.Time)
				//Prints message
				fmt.Printf("T:%d - %v -> %v \n", myClock.GetTime(), in.Id, in.Message)
			}
		}
	}()

	<-waitc
}

func leave(ctx context.Context, client pb.ServiceClient) {
	myClock.Increment()

	//Logs the request to leave
	log.Printf("Client %v is trying to leave", MyId)

	//Creates leave request
	leaveRequest := &pb.LeaveRequest{
		Id:   MyId,
		Time: myClock.GetTime(),
	}

	//Calls the Leave method defined in the proto file
	reponse, err := client.Leave(ctx, leaveRequest)
	myClock.SyncTime(reponse.Time)
	if err != nil {
		log.Printf("T:%d - Leave failed - %v", myClock.GetTime(), reponse.HttpStatusCode)
	} else {
		log.Printf("T:%d - Leave successful - %v", myClock.GetTime(), reponse.HttpStatusCode)
	}
}

//Creates the ID of the client by calling to the server and getting the next id number available
func createID(client pb.ServiceClient, ctx context.Context) {
	tempid, err := client.CreateID(ctx, &pb.NodeId{Id: MyId})
	if err != nil {
		log.Printf("Failed to get id")
	}
	MyId = tempid.Id
}
