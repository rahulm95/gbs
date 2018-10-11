package models

import (
	"context"
	"log"

	pb "github.com/voonik/goConnect/api/go/groupbuy"
	"github.com/voonik/goFramework/pkg/database"
)

// GameGroup ...
type GameGroup struct {
	database.VModel
	CreatedBy       uint64
	Status          status
	RuleID          uint64
	UUID            string
	GameGroupOrders []GameGroupOrder `gorm:"many2many:game_group_orders;"`
}

// GameGroupInterface ...
type GameGroupInterface interface {
	GetGroupCount(pb.GroupRequest) pb.GroupCount
	isGroupOpen(context.Context, pb.GroupJoin) bool
}

type status uint

// List of status for Group
const (
	OPENGROUP status = 1 + iota
	CLOSEDGROUP
	EXPIREDGROUP
	PENDINGGROUP
)

// GetGroupCount returns the number of groups for a product respective to the status of the group
func (gg GameGroup) GetGroupCount(gp pb.GroupRequest) pb.GroupCount {
	log.Println("Getting group count")
	// condition := "productID = " + strconv.FormatInt(gp.ProductID, 16) + string(gp.Status)
	// if gp.Status != 0 {
	// 	condition += "and status = " + strconv.FormatInt(int64(gp.Status), 16)
	// }
	var gc pb.GroupCount
	return gc
}

// IsGroupOpen ...
func (gg GameGroup) IsGroupOpen(ctx context.Context, group *pb.GroupJoin) bool {
	log.Println("Getting group status")
	gg.ID = group.GroupID
	gg.Status = OPENGROUP
	var count int
	DB(ctx).Table("game_groups").Where(&gg).Count(&count)
	log.Println("value", count)
	if count == 1 {
		return true
	} else {
		return false
	}
}

// CreateGameGroup ...
func (gg GameGroup) CreateGameGroup(ctx context.Context, group *pb.Group) pb.Response {
	log.Println("Creating the new group")
	newGroup := &GameGroup{
		CreatedBy: group.CreatedBy,
		Status:    CLOSEDGROUP,
	}
	err := DB(ctx).NewRecord(&newGroup)
	log.Println(err)
	var res pb.Response
	res.ResponseText = "success"
	res.ResponseCode = 200
	return res
}

// ListGroups ...
func (gg GameGroup) ListGroups(ctx context.Context, groupRequest *pb.GroupRequest) (groups pb.Groups) {
	log.Println("Listing groups for a product")
	// ggg := GameGroup{}
	ctx = inten.SetInContextThreadObject(ctx, nil)
	DB(ctx).Model(&gg).Joins("join rules on rules.id = game_groups.rule_id and rules.rule_type='nuserbuymproduct'").Joins("join game_group_orders on game_group_orders.game_group_id = game_groups.id and game_group_orders.product_id = rules.param_2").Where(&GameGroup{Status: 1}).Where(&GameGroupOrder{ProductID: (groupRequest.ProductID)}).Find(&gg)
	log.Println("values", &gg)
	DB(ctx).Model(&gg).Joins("join rules on rules.id = game_groups.rule_id and rules.rule_type='nuserbuyinmcollection'").Joins("join game_group_orders on game_group_orders.game_group_id = game_groups.id").Joins("JOIN products_group_product_collections on products_group_product_collections.product_id = game_group_orders.product_id").Joins("JOIN group_product_collections on group_product_collections.id = products_group_product_collections.group_product_collection_id and is_active =1").Where(&GameGroup{Status: 1}).Where(&GameGroupOrder{ProductID: (groupRequest.ProductID)}).Find(&gg)
	log.Println("values", &gg)
	return groups
}
