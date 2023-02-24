package main

import (
	"fmt"

	"group-logic-v0.1/p2p"
	"group-logic-v0.1/pubsub"
)

const groupName string = "test"
const key string = "dog-cat-both-are"

func main() {
	privateGroupKey := groupName + "/" + key
	ctx, host := p2p.EstablishP2P()
	grp := pubsub.HandlePubSub(ctx, host, groupName, privateGroupKey)
	fmt.Println(grp.PrivateGroupSub)
	fmt.Println(grp.PrivateGroupTopic)
	fmt.Println(grp.PublicGroupSub)
	fmt.Println(grp.PublicGroupTopic)
	for {

	}
}
