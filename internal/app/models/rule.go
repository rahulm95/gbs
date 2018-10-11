package models

import (
	"context"
	"log"

	pb "github.com/voonik/goConnect/api/go/groupbuy"
	"github.com/voonik/goFramework/pkg/database"
)

// RuleInterface specifies how the group condition is met
type RuleInterface interface {
	ValidateOrder(context.Context, pb.GroupID) bool
}

// Rule specifies how the group condition is met
type Rule struct {
	database.VModel
	RuleType string
}

// List of rules
const (
	NUSERBUYMPRODUCT      = "nuserbuymproduct"
	NUSERBUYINMCOLLECTION = "nuserbuyinmcollection"
)

type ruleValidatorInterface interface {
	ValidateOrder(context.Context, uint64, uint64) bool
}

var ruleRegister = make(map[string]ruleValidatorInterface)

// RegisterRule all available rules
func RegisterRule(rule string, ruleInstance ruleValidatorInterface) {
	log.Println("Registering rule", rule)
	ruleRegister[rule] = ruleInstance
}

// ValidateOrder ...
func (rule Rule) ValidateOrder(ctx context.Context, groupID *pb.GroupID) bool {
	gg := &GameGroup{}
	DB(ctx).Select("rule_id").Find(&gg, "id = ?", groupID.GroupID)
	log.Println("value of rule ID", gg.RuleID)
	ruleObject := &Rule{}
	DB(ctx).Select("rule_type").First(&ruleObject, gg.RuleID)
	log.Println("value of rule type", ruleObject.RuleType)
	ruleValidator := ruleRegister[ruleObject.RuleType]
	log.Println(ruleValidator)
	return ruleValidator.ValidateOrder(ctx, groupID.GroupID, gg.RuleID)
}
