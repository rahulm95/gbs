package services

import (
	"context"
	"log"

	"github.com/voonik/gbs/internal/app/models"
	pb "github.com/voonik/goConnect/api/go/groupbuy"
)

// GameGroupOrderService ...
type GameGroupOrderService struct {
	gameGroupOrder models.GameGroupOrder
	gameGroup      models.GameGroup
	rule           models.Rule
}

// GameGroupOrderServiceInterface ...
type GameGroupOrderServiceInterface interface {
	JoinGroupOrder(context.Context, *pb.GroupJoin) (pb.Response, error)
	CancelGroupOrder(context.Context, *pb.Group) (pb.Response, error)
}

//CancelGroupOrder ...
func (ggo *GameGroupOrderService) CancelGroupOrder(ctx context.Context, orderDetail *pb.Order) (*pb.Response, error) {
	shouldReconcileOrder, resp := ggo.gameGroupOrder.CancelGroupOrder(ctx, orderDetail)
	resp.ResponseCode = 200
	if shouldReconcileOrder {
		resp.ResponseText = "The group order has been cancelled"
	}
	resp.ResponseText = "The group order not found"
	var err error
	return &resp, err
}

//JoinGroupOrder ...
func (ggo *GameGroupOrderService) JoinGroupOrder(ctx context.Context, group *pb.GroupJoin) (*pb.Response, error) {
	var resp pb.Response
	var err error
	if ggo.gameGroup.IsGroupOpen(ctx, group) {
		ggo.gameGroupOrder.JoinGroupOrder(ctx, group)
		// ggo.rule.ValidateOrder(ctx,)
		resp.ResponseText = "Joined group successfully"
	} else {
		resp.ResponseText = "The group has been completed already"
	}
	log.Println("resp", resp)
	return &resp, err
}
