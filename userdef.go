package main

import (
	"fmt"
	"net"
	"net/rpc"
	"strconv"
	"test/network/chord"
)

func NewNode(port int) dhtNode {
	newClient := new(chord.Client)
	newClient.Node_.Ip = ":" + strconv.Itoa(port)
	newClient.Node_.Predecessor = nil

	newClient.Server = rpc.NewServer()
	err := newClient.Server.Register(&newClient.Node_)
	if err != nil {
		fmt.Println(err)
	}

	listener, e := net.Listen("tcp", newClient.Node_.Ip)
	newClient.Listener = listener
	if e != nil {
		fmt.Println(e)
	}
	go newClient.Server.Accept(newClient.Listener)
	newClient.Node_.Listening = true
	newClient.Node_.Ip = chord.GetLocalAddress() + newClient.Node_.Ip
	newClient.Node_.Id = chord.HashString(newClient.Node_.Ip)
	var res = newClient
	return res
}

//func NewNode(port int) *kademlia.Client {
//	newClient := new(kademlia.Client)
//	newClient.Node_.RoutingTable.Ip = ":" + strconv.Itoa(port)
//	newClient.Server = rpc.NewServer()
//	err := newClient.Server.Register(&newClient.Node_)
//	if err != nil {
//		fmt.Println(err)
//	}
//	return newClient
//}
