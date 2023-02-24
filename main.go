package main

import (
	"fmt"

	"group-logic-v0.1/p2p"
	"group-logic-v0.1/p2p/peerdisc"
	"group-logic-v0.1/pubsub"
	"group-logic-v0.1/pubsub/msghandle"
)

const groupName string = "test"
const key string = "dog-cat-both-are"
const service string = "service/test/rex"

func main() {
	privateGroupKey := groupName + "/" + key
	ctx, host := p2p.EstablishP2P()
	grp := pubsub.HandlePubSub(ctx, host, groupName, privateGroupKey)
	fmt.Println(grp.PrivateGroupSub)
	fmt.Println(grp.PrivateGroupTopic)
	fmt.Println(grp.PublicGroupSub)
	fmt.Println(grp.PublicGroupTopic)
	kadDHT := p2p.HandleDHT(ctx, host)
	go peerdisc.DiscoverPeers(ctx, host, kadDHT, service)
	msghandle.HandlePubSubMessages(ctx, host, grp.PrivateGroupSub, grp.PrivateGroupTopic)
	msghandle.HandlePubSubMessages(ctx, host, grp.PublicGroupSub, grp.PublicGroupTopic)
	for {

	}
}
