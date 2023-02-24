package p2p

import (
	"context"
	"fmt"

	libp2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
)

func EstablishP2P() (context.Context, host.Host) {
	nat := libp2p.NATPortMap()
	holePunch := libp2p.EnableHolePunching()
	transPort := libp2p.Transport(tcp.NewTCPTransport)
	host, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"), nat, holePunch, transPort)
	if err != nil {
		fmt.Println("Error during  starting the node")
	}
	ctx := context.Background()
	return ctx, host

}
