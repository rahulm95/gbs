package main

import (
	"context"
	"log"

	"github.com/voonik/gbs/internal/app/handlers"
	"github.com/voonik/gbs/internal/app/services"
	pb "github.com/voonik/goConnect/api/go/groupbuy"
	"github.com/voonik/goFramework/pkg/grpc/server"
)

func main() {
	log.Println("Setting up Framework")
	server.Init()
	// pb.RegisterGameGroupServer(server.GrpcServer, handlers.GetGameGroupOrderInstance())
	gg := handlers.GetGameGroupInstance()

	// ggo := handlers.GetGameGroupOrderInstance()
	// gb := &groupbuy.Group{
	// 	UserID:  123,
	// 	GameID:  2,
	// 	OrderID: 33123,
	// }

	// groupID := pb.GroupID{}
	// groupID.GroupID = 16
	// gg.CreateGameGroup(context.Background(), gb)
	// gg.GetGroupStatus(context.Background(), &groupID)

	listGroups(gg)
	// cancelOrder(ggo)
	// joinOrder(ggo)
	// defer server.Finish()

}

func joinOrder(ggo *services.GameGroupOrderService) {
	gb := &pb.GroupJoin{
		UserID: 932382,
		// GroupID:   16,
		OrderID: 9121,
		// ProductID: 876721,
	}
	ggo.JoinGroupOrder(context.Background(), gb)
}

func cancelOrder(ggo *services.GameGroupOrderService) {
	orderDetails := pb.Order{
		OrderID: 671231,
	}
	ggo.CancelGroupOrder(context.Background(), &orderDetails)
}

func listGroups(gg *services.GameGroupService) {
	groupRequest := &pb.GroupRequest{
		Status:    1,
		ProductID: 500001,
	}
	gg.ListGroups(context.Background(), groupRequest)
}
