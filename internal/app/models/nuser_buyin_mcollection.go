package models

import (
	"context"
	"log"
)

// NuserBuyinMcollection ...
type NuserBuyinMcollection struct {
	Users                    uint16 `gorm:"column:param_1"`
	GroupProductCollectionID uint64 `gorm:"column:param_2"`
}

// NuserBuyinMcollectionInterface ...
type NuserBuyinMcollectionInterface interface {
}

// TableName ...
func (NuserBuyinMcollection) TableName() string {
	return "rules"
}

func init() {
	RegisterRule(NUSERBUYINMCOLLECTION, &NuserBuyinMcollection{})
}

// ValidateOrder ...
func (rule *NuserBuyinMcollection) ValidateOrder(ctx context.Context, groupID uint64, ruleID uint64) bool {
	ruleObject := rule.GetParameters(ctx, ruleID)
	return rule.RuleCalculator(ctx, ruleObject, groupID)
}

// GetParameters ...
func (rule *NuserBuyinMcollection) GetParameters(ctx context.Context, ruleID uint64) *NuserBuyinMcollection {
	// DB(ctx).Model(&NuserBuyinMcollection{}).Find()
	ruleObject := &NuserBuyinMcollection{}
	DB(ctx).Find(&ruleObject, "id = ?", ruleID)
	log.Println("ruleid", ruleObject)
	return ruleObject
	// DB(ctx).Create(&ruleObject)
}

// RuleCalculator ...
func (rule *NuserBuyinMcollection) RuleCalculator(ctx context.Context, ruleObject *NuserBuyinMcollection, groupID uint64) bool {
	log.Printf("inside NUSERBUYINMCOLLECTION calculator")
	ggo := &GameGroupOrder{}
	ggo.GameGroupID = groupID
	var totalCount int
	DB(ctx).Model(&ggo).Joins("JOIN products_group_product_collections on products_group_product_collections.product_id = game_group_orders.product_id and products_group_product_collections.group_product_collection_id = ?", ruleObject.GroupProductCollectionID).Joins("JOIN group_product_collections on group_product_collections.id = products_group_product_collections.group_product_collection_id and is_active =1").Where(&GameGroupOrder{GameGroupID: groupID}).Count(&totalCount)
	log.Println("ggo", totalCount)
	groupStatus := uint16(totalCount) == ruleObject.Users
	return groupStatus
}
