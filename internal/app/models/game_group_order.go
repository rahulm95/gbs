package models

import (
	"context"
	"log"

	"github.com/jinzhu/gorm"
	pb "github.com/voonik/goConnect/api/go/groupbuy"
)

// GameGroupOrder stores game to order mapping
type GameGroupOrder struct {
	gorm.Model
	GameGroupID uint64
	OrderID     uint64
	ProductID   uint64
	UserID      uint64
	OrderStatus status
}

// List of status for GroupOrder
const (
	OPENGROUPORDER status = 1 + iota
	CANCELLEDGROUPORDER
)

// GameGroupOrderInterface ...
type GameGroupOrderInterface interface {
	JoinGroupOrder(context.Context, *pb.Group) (pb.Response, error)
	CancelGroupOrder(context.Context, *pb.Order) (bool, pb.Response)
}

// JoinGroupOrder ...
func (ggo GameGroupOrder) JoinGroupOrder(ctx context.Context, group *pb.GroupJoin) pb.Response {
	ggo.OrderID = group.OrderID
	ggo.GameGroupID = group.GroupID
	ggo.UserID = group.UserID
	ggo.ProductID = group.ProductID
	ggo.OrderStatus = OPENGROUPORDER
	DB(ctx).Create(&ggo)
	var res pb.Response
	res.ResponseText = "success"
	res.ResponseCode = 200
	return res
}

// CancelGroupOrder ...
func (ggo GameGroupOrder) CancelGroupOrder(ctx context.Context, orderDetail *pb.Order) (bool, pb.Response) {
	log.Println("Cancelling one of the group orders:", orderDetail.OrderID)
	DB(ctx).Model(&ggo).Where("order_id = ?", orderDetail.OrderID).Find(&ggo)
	shouldReconcileOrder := false
	var resp pb.Response
	if ggo.OrderID != 0 {
		rowsAffected := DB(ctx).Table("game_group_orders").Where("order_id = ?", ggo.OrderID).Update("order_status", 2).RowsAffected
		if rowsAffected > 0 {
			shouldReconcileOrder = true
			resp.ResponseText = "The group order has been cancelled"
		} else {
			resp.ResponseText = "The group order has already been cancelled"
		}
	} else {
		resp.ResponseText = "Group Order not found"
	}
	log.Println("values before return", shouldReconcileOrder, resp)
	return shouldReconcileOrder, resp
}
