package pubsub

import (
	"context"
	"fmt"

	pbsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
)

type groupService struct {
	privateGroup *pbsub.PubSub
	exposeGroup  *pbsub.PubSub
}

type Group struct {
	PrivateGroupTopic *pbsub.Topic
	PrivateGroupSub   *pbsub.Subscription
	PublicGroupTopic  *pbsub.Topic
	PublicGroupSub    *pbsub.Subscription
}

func newGroupService(ctx context.Context, host host.Host) *groupService {
	internalPubsub, err := pbsub.NewGossipSub(ctx, host)
	if err != nil {
		fmt.Println("Error during creating the group private service")
	}
	exposedPubsub, err := pbsub.NewGossipSub(ctx, host)
	if err != nil {
		fmt.Println("Error during creation of the exposed group service")
	}
	grpSrv := &groupService{
		privateGroup: internalPubsub,
		exposeGroup:  exposedPubsub,
	}
	return grpSrv
}

func HandlePubSub(ctx context.Context, host host.Host, topic string, topicprv string) *Group {
	grpSrv := newGroupService(ctx, host)
	prvTopic, err := grpSrv.privateGroup.Join(topicprv)
	if err != nil {
		fmt.Println("Error while joining the private topic")
	}
	prvSub, err := grpSrv.privateGroup.Subscribe(topicprv)
	if err != nil {
		fmt.Println("Error while subscribing to the private topic")
	}
	expTopic, err := grpSrv.exposeGroup.Join(topic)
	if err != nil {
		fmt.Println("Error while joining the public topic")
	}
	expSub, err := grpSrv.privateGroup.Subscribe(topic)
	if err != nil {
		fmt.Println("Error while subscribing to the public topic")
	}

	return &Group{
		PrivateGroupTopic: prvTopic,
		PrivateGroupSub:   prvSub,
		PublicGroupTopic:  expTopic,
		PublicGroupSub:    expSub,
	}

}
