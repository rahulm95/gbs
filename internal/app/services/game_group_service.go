package services

import (
	"context"

	"github.com/voonik/gbs/internal/app/models"
	pb "github.com/voonik/goConnect/api/go/groupbuy"
	"github.com/voonik/goFramework/pkg/logger"
)

// GameGroupService invokes the service methods
type GameGroupService struct {
	gameGroup models.GameGroup
	Rule      models.Rule
}

// GameGroupServiceInterface ...
type GameGroupServiceInterface interface {
	CreateGameGroup(context.Context, pb.Group)
}

//CreateGameGroup ...
func (gg *GameGroupService) CreateGameGroup(ctx context.Context, gb *pb.Group) (response *pb.Response, err error) {
	logger.Print(gb)
	gg.gameGroup.CreateGameGroup(ctx, gb)
	return response, err
}

//GetGameGroupDetails ...
func (gg *GameGroupService) GetGameGroupDetails(context.Context, *pb.Group) (response *pb.Response, err error) {

	return response, err
}

//ListGroups ...
func (gg *GameGroupService) ListGroups(ctx context.Context, groupRequest *pb.GroupRequest) (groups *pb.Groups, err error) {

	gg.gameGroup.ListGroups(ctx, groupRequest)
	return groups, err
}

//GetGroupCount ...
func (gg *GameGroupService) GetGroupCount(context.Context, *pb.GroupRequest) (groupCount *pb.GroupCount, err error) {

	return groupCount, err
}

//GetGroupStatus ...
func (gg *GameGroupService) GetGroupStatus(ctx context.Context, groupID *pb.GroupID) (groupStatus *pb.GroupStatus, err error) {
	// gg.rule.ValidateOrder(ctx, groupID)
	return groupStatus, err
}
