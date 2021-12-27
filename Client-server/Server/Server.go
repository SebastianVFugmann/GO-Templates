package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"sync"

	//The link to the folder that the package is in, but still from github
	pb "github.com/SebastianVFugmann/GO-Templates/Client-server/Proto"
	"github.com/SebastianVFugmann/GO-Templates/LamportClock"
	"google.golang.org/grpc"
)

const (
	port = ":6969"
)

func main() {
	//Checks of a log file exists and then appends or creates a new
	file, logerr := os.OpenFile("server_logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if logerr != nil {
		log.Fatal(logerr)
	}
	//Sets the loggers output to the log file
	log.SetOutput(file)

	// Creates listener instance
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Creates server instance
	s := grpc.NewServer()

	//Registers chitty chat
	cc := CreateServer()
	pb.RegisterServiceServer(s, cc)

	// Listen and serve
	log.Printf("T:%d - server listening at %v \n", cc.lamport.GetTime(), lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

//The server and its variables
type Server struct {
	pb.UnimplementedServiceServer
	Clients map[int32]chan *pb.Object
	lamport LamportClock.LamportClock
	mu      sync.RWMutex
	ids     int32
}

//Creates the server by initializing the client map and the lamport clock
func CreateServer() *Server {
	server := &Server{
		Clients: make(map[int32]chan *pb.Object),
	}
	server.lamport.Initialize()
	server.ids = 0
	return server
}

//Used by a new client to join the chat service
//Also one of the methods defined by the proto file
func (s *Server) Join(in *pb.JoinRequest, msgStream pb.Service_JoinServer) error {
	s.lamport.SyncTime(in.Time)
	log.Printf("T:%d - Received a JOIN message from client %d \n", s.lamport.GetTime(), in.Id)

	//Adds the client to the map of clients
	err := s.addClient(in.Id)
	if err != nil {
		log.Printf("Error in joining with client: %v, error: %v", in.Id, err)
	}

	// doing this never closes the stream
	for {
		select {
		case <-msgStream.Context().Done():
			return nil
		case msg := <-s.Clients[in.Id]:
			log.Printf("T:%d - Client: %v got message: %v \n", s.lamport.GetTime(), in.Id, msg)
			msgStream.Send(msg)
		}
	}
}

//Method that adds the client to the map of clients
func (s *Server) addClient(id int32) error {
	//Increments clock
	s.lamport.Increment()
	s.mu.Lock()
	defer s.mu.Unlock()

	// Creates new client and checks if the id exists or is empty
	_, ok := s.Clients[id]
	if ok {
		return errors.New("client has already joined, cant rejoin")
	}

	// Creates a new channel for the new client
	s.Clients[id] = make(chan *pb.Object)

	log.Printf("T:%d - Added new client %d", s.lamport.GetTime(), id)
	return nil
}

func (s *Server) Leave(ctx context.Context, in *pb.LeaveRequest) (*pb.ServerResponse, error) {
	s.lamport.SyncTime(in.Time)

	//Removes the client from the server's map
	delete(s.Clients, in.Id)

	return &pb.ServerResponse{
		HttpStatusCode: 200,
		Time:           s.lamport.GetTime(),
	}, nil
}

//Gets the next ID and sends it to the client
func (server *Server) CreateID(ctx context.Context, in *pb.NodeId) (*pb.NodeId, error) {
	server.ids += 1
	return &pb.NodeId{
		Id: server.ids,
	}, nil
}
