package models

import (
	"context"
	"log"
)

// NuserBuyMproduct ...
type NuserBuyMproduct struct {
	Users     uint16 `gorm:"column:param_1"`
	ProductID uint64 `gorm:"column:param_2"`
}

// NuserBuyMproductInterface ...
type NuserBuyMproductInterface interface {
}

// TableName ...
func (NuserBuyMproduct) TableName() string {
	return "rules"
}

func init() {
	RegisterRule(NUSERBUYMPRODUCT, &NuserBuyMproduct{})
}

// ValidateOrder ...
func (rule *NuserBuyMproduct) ValidateOrder(ctx context.Context, groupID uint64, ruleID uint64) bool {
	ruleObject := rule.GetParameters(ctx, ruleID)
	satisfied := rule.RuleCalculator(ctx, ruleObject, groupID)
	return satisfied
}

// GetParameters ...
func (rule *NuserBuyMproduct) GetParameters(ctx context.Context, ruleID uint64) *NuserBuyMproduct {
	// DB(ctx).Model(&NuserBuyMproduct{}).Find()
	ruleObject := &NuserBuyMproduct{}
	DB(ctx).Find(&ruleObject, "id = ?", ruleID)
	log.Println(ruleObject)
	return ruleObject
	// DB(ctx).Create(&ruleObject)
}

// RuleCalculator ...
func (rule *NuserBuyMproduct) RuleCalculator(ctx context.Context, ruleObject *NuserBuyMproduct, groupID uint64) bool {
	ggo := &GameGroupOrder{}
	var count int
	DB(ctx).Model(&ggo).Where("game_group_id = ? and product_id = ?", groupID, ruleObject.ProductID).Count(&count)
	log.Println("count of orders", count)
	return false
}
