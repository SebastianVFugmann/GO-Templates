package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	pb "github.com/SebastianVFugmann/GO-Templates/Active_replication/Service"
)

type frontend struct {
	id      int32
	clients map[int32]pb.AuctionClient
	ctx     context.Context
	ackch   chan pb.Acknowledgement
	repch   chan pb.StatusReply
}

func (fe *frontend) bidOnReplica(m *pb.BidMessage, c pb.AuctionClient) {
	ack, err := c.Bid(fe.ctx, m)
	if err != nil {
		if strings.Contains(err.Error(), "Not a valid auction id.") {
			fe.ackch <- pb.Acknowledgement{Ack: -1}
		} else {
			fe.ackch <- pb.Acknowledgement{Ack: -2}
		}
	} else {
		fe.ackch <- *ack
	}
}

func (fe *frontend) bid(auctionId int32, bid int32) {
	m := &pb.BidMessage{
		AuctionId: auctionId,
		Bid:       bid,
		Bidder:    fe.id,
	}
	var a pb.Acknowledgement
	var r pb.StatusReply

bidonallreplicas:
	for k, v := range fe.clients {
		go fe.bidOnReplica(m, v)

		start := time.Now()
	fivesecondcheck:
		for start.Add(5 * time.Second).After(time.Now()) {

			select {
			case ack := <-fe.ackch:
				if ack.Ack == -1 {
					fmt.Printf("This is not a valid auction id: %v.\n", auctionId)
					return
				} else if ack.Ack == -2 {
					// outcommented because of !transparency
					// log.Println("Server seems to be dead. Deleting ... ")
					delete(fe.clients, k)
					continue bidonallreplicas
				}
				a = ack
				break fivesecondcheck
			default:
			}
		}

		req := &pb.StatusRequest{
			AuctionId: auctionId,
		}
		reply, _ := v.Status(fe.ctx, req)
		r = *reply
	}

	switch acktype := a.Ack; acktype {
	case 0:
		fmt.Printf("Bid of %v camels accepted. You are currently the highest bidder on auction %v.\n", bid, auctionId)
	case 1:
		if r.HighestBid == 0 {
			fmt.Print("You must bid at least 1 camel.\n")
		} else {
			fmt.Printf("Bid of %v camels not accepted. There is a higher bid of %v camels from %v on auction %v.\n", bid, r.HighestBid, r.Bidder, auctionId)
		}
	case 2:
		if r.HighestBid == 0 {
			fmt.Printf("Bid of %v camels not accepted. Auction %v has ended. No bids were made.\n", bid, auctionId)
		} else {
			fmt.Printf("Bid of %v camels not accepted. Auction %v has ended. Winner: %v with a bid of %v camels.\n", bid, auctionId, r.Bidder, r.HighestBid)
		}
	}
}

func (fe *frontend) statusFromReplica(req *pb.StatusRequest, c pb.AuctionClient) {
	rep, err := c.Status(fe.ctx, req)
	if err != nil {
		if strings.Contains(err.Error(), "Not a valid auction id.") {
			fe.repch <- pb.StatusReply{HighestBid: -1}
		} else {
			fe.repch <- pb.StatusReply{HighestBid: -2}
		}
	} else {
		fe.repch <- *rep
	}
}

func (fe *frontend) status(auctionId int32) {
	req := &pb.StatusRequest{
		AuctionId: auctionId,
	}
	var r pb.StatusReply

statusfromallreplicas:
	for k, v := range fe.clients {
		go fe.statusFromReplica(req, v)

		start := time.Now()
	fivesecondcheck:
		for start.Add(5 * time.Second).After(time.Now()) {
			select {
			case rep := <-fe.repch:
				if rep.HighestBid == -1 {
					fmt.Printf("This is not a valid auction id: %v.\n", auctionId)
					return
				} else if rep.HighestBid == -2 {
					// log.Println("Server seems to be dead. Deleting ... ")
					delete(fe.clients, k)
					continue statusfromallreplicas
				}
				r = rep
				break fivesecondcheck
			default:
			}
		}
	}

	if r.Active {
		if r.HighestBid == 0 {
			fmt.Print("This auction is still active. No bids have been made yet.\n")
		} else {
			fmt.Printf("This auction is still active. Highest bid is currently %v camels from %v.\n", r.HighestBid, r.Bidder)
		}
	} else {
		if r.HighestBid == 0 {
			fmt.Print("This auction has ended. No bids were made.\n")
		} else {
			fmt.Printf("This auction has ended. Winning bid was %v camels from %v.\n", r.HighestBid, r.Bidder)
		}
	}
}
