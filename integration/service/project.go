package service

import (
	"encoding/json"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	integrationDomain "github.com/aaronchen2k/deeptest/integration/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
)

type ProjectService struct {
	RemoteService   *RemoteService        `inject:""`
	IntegrationRepo *repo.IntegrationRepo `inject:""`
	ProjectRepo     *repo.ProjectRepo     `inject:""`
	UserRepo        *repo.UserRepo        `inject:""`
	ProjectRoleRepo *repo.ProjectRoleRepo `inject:""`
	MessageRepo     *repo.MessageRepo     `inject:""`
	BaseRepo        *repo.BaseRepo        `inject:""`
	MessageService  *MessageService       `inject:""`
}

func (s *ProjectService) GetUserProductList(page, pageSize int, username string) (ret []integrationDomain.ProductItem, err error) {
	return s.RemoteService.GetUserProductList(page, pageSize, username)
}

func (s *ProjectService) GetSpacesByUsername(username string) (ret []integrationDomain.SpaceItem, err error) {
	return s.RemoteService.GetSpacesByUsername(username)
}

func (s *ProjectService) AddProjectRelatedProducts(projectId uint, products []uint) (err error) {
	relations := make([]model.ProjectProductRel, 0)
	for _, product := range products {
		relations = append(relations, model.ProjectProductRel{
			ProjectId: projectId,
			ProductId: product,
		})
	}

	if len(relations) > 0 {
		err = s.IntegrationRepo.BatchCreateProjectProductRel(relations)
	}

	return
}

func (s *ProjectService) AddProjectRelatedSpaces(projectId uint, spaces []string) (err error) {
	relations := make([]model.ProjectSpaceRel, 0)
	for _, space := range spaces {
		relations = append(relations, model.ProjectSpaceRel{
			ProjectId: projectId,
			SpaceCode: space,
		})
	}

	if len(relations) > 0 {
		err = s.IntegrationRepo.BatchCreateProjectSpaceRel(relations)
	}

	return
}

func (s *ProjectService) SyncSpaceMembers(projectId uint, spaces []string) (err error) {
	members, memberRoles, err := s.GetUserInfoMap(spaces)
	if err != nil {
		return
	}

	err = s.AddMembers(projectId, members, memberRoles)

	return
}

func (s *ProjectService) Save(req integrationDomain.ProjectReq, projectId uint) (err error) {
	if config.CONFIG.System.SysEnv != "ly" {
		return
	}

	if err = s.SaveProducts(projectId, req.Products); err != nil {
		return
	}

	if err = s.SaveSpaces(projectId, req.Spaces); err != nil {
		return
	}

	if req.SyncMembers {
		err = s.SyncSpaceMembers(projectId, req.Spaces)
	}

	return
}

func (s *ProjectService) SaveProducts(projectId uint, products []uint) (err error) {
	if err = s.IntegrationRepo.DeleteProductByProject(projectId); err != nil {
		return
	}

	err = s.AddProjectRelatedProducts(projectId, products)

	return
}

func (s *ProjectService) SaveSpaces(projectId uint, spaces []string) (err error) {
	if err = s.IntegrationRepo.DeleteSpaceByProject(projectId); err != nil {
		return
	}

	err = s.AddProjectRelatedSpaces(projectId, spaces)

	return
}

func (s *ProjectService) GetProductsByProject(projectId uint) (res []integrationDomain.ProductBaseItem, err error) {
	productIds, err := s.IntegrationRepo.GetProductsByProject(projectId)
	if err != nil {
		return
	}

	res, err = s.RemoteService.GetProductListById(productIds)

	return
}

func (s *ProjectService) GetSpacesByProject(projectId uint) (res []integrationDomain.SpaceItem, err error) {
	spaceCodes, err := s.IntegrationRepo.GetSpacesByProject(projectId)
	if err != nil {
		return
	}

	res, err = s.RemoteService.BatchGetSpacesByCode(spaceCodes)

	return
}

func (s *ProjectService) Detail(projectId uint) (res integrationDomain.ProjectDetail, err error) {
	products, err := s.GetProductsByProject(projectId)
	if err != nil {
		return
	}

	spaces, err := s.GetSpacesByProject(projectId)
	if err != nil {
		return
	}

	res.Products = products
	res.Spaces = spaces

	return
}

func (s *ProjectService) GetListWithRoleBySpace(spaceCode, username string) (res []v1.ProjectListWithRole, err error) {
	res, err = s.IntegrationRepo.GetProjectListWithRoleBySpace(spaceCode)
	if err != nil {
		return
	}

	s.AddMemberRoleForProject(&res, username)

	return
}

func (s *ProjectService) AddMemberRoleForProject(projects *[]v1.ProjectListWithRole, username string) {
	projectIds := make([]uint, 0)
	for _, v := range *projects {
		projectIds = append(projectIds, v.ID)
	}

	projectRoleMap, err := s.ProjectRepo.GetUserProjectRoleMap(username, projectIds)
	if err != nil {
		return
	}

	for _, v := range *projects {
		if roleName, ok := projectRoleMap[v.ID]; ok {
			v.RoleName = roleName
		}
	}

	return
}

func (s *ProjectService) SaveSpaceRelatedProjects(spaceCode string, shortNames []string) (err error) {
	err = s.IntegrationRepo.DeleteBySpaceCode(spaceCode)
	if err != nil {
		return
	}

	projectShortNameIdMap, err := s.GetProjectShortNameAndIdMap(shortNames)
	if err != nil {
		return
	}

	relations := make([]model.ProjectSpaceRel, 0)
	for _, shortName := range shortNames {
		relTmp := model.ProjectSpaceRel{
			SpaceCode: spaceCode,
		}
		if projectId, ok := projectShortNameIdMap[shortName]; ok {
			relTmp.ProjectId = projectId
		}

		relations = append(relations, relTmp)
	}

	if len(relations) > 0 {
		err = s.IntegrationRepo.BatchCreateProjectSpaceRel(relations)
	}
	
	return
}

func (s *ProjectService) GetProjectShortNameAndIdMap(shortNames []string) (res map[string]uint, err error) {
	projects, err := s.ProjectRepo.BatchGetByShortNames(shortNames)
	if err != nil {
		return
	}

	res = make(map[string]uint)
	for _, project := range projects {
		res[project.ShortName] = project.ID
	}

	return
}

func (s *ProjectService) GetUserInfoMap(spaceCodes []string) (res map[string]integrationDomain.UserRoleInfo, userRoles map[string][]string, err error) {
	spaceMemberRoles, err := s.RemoteService.BatchGetMembersBySpaces(spaceCodes)
	if err != nil {
		return
	}

	res = make(map[string]integrationDomain.UserRoleInfo)
	userRoles = make(map[string][]string)

	for _, memberRoles := range spaceMemberRoles {
		for _, memberRole := range memberRoles.UserBaseInfo {
			res[memberRole.Username] = memberRole
			for _, role := range memberRole.Role {
				userRoles[memberRole.Username] = append(userRoles[memberRole.Username], role.Value)
			}
		}
	}

	roleBased, err := s.GetSpaceRoleArrays()
	if err != nil {
		return
	}

	for k, v := range userRoles {
		roles := _commUtils.ArrayRemoveDuplication(v)
		userRoles[k] = _commUtils.Intersect(roles, roleBased)
	}

	return
}

func (s *ProjectService) GetSpaceRoleArrays() (res []string, err error) {
	spaceRoles, err := s.RemoteService.GetSpaceRoles()
	if err != nil {
		return
	}

	for _, v := range spaceRoles {
		res = append(res, v.RoleValue)
	}

	return
}

func (s *ProjectService) AddMembers(projectId uint, members map[string]integrationDomain.UserRoleInfo, memberRoles map[string][]string) (err error) {
	for _, member := range members {
		createUserReq := v1.UserReq{
			UserBase: v1.UserBase{
				Username:  member.Username,
				Name:      member.RealName,
				Email:     member.Mail,
				ImAccount: member.WxName,
				Password:  _commUtils.RandStr(8),
			},
		}
		userId, err := s.UserRepo.CreateIfNotExisted(createUserReq)
		if err != nil {
			continue
		}

		role := consts.IntegrationGeneral
		if roles, ok := memberRoles[member.Username]; ok && len(roles) > 0 {
			role = consts.RoleType(roles[0])
		}

		err = s.ProjectRepo.AddMemberIfNotExisted(projectId, userId, role)
	}

	return
}

func (s *ProjectService) SendApplyMessage(projectId, userId, auditId uint, roleName consts.RoleType) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("发送消息异常")
		}
	}()

	if config.CONFIG.System.SysEnv != "ly" {
		return
	}

	messageContent, err := s.MessageService.GetJoinProjectMcsData(userId, projectId, auditId, roleName)
	messageContentByte, _ := json.Marshal(messageContent)

	adminRole, err := s.ProjectRoleRepo.FindByName(s.BaseRepo.GetAdminRoleName())
	if err != nil {
		return
	}

	messageReq := v1.MessageReq{
		MessageBase: v1.MessageBase{
			MessageSource: consts.MessageSourceJoinProject,
			Content:       string(messageContentByte),
			ReceiverRange: 3,
			SenderId:      userId,
			ReceiverId:    adminRole.ID,
			SendStatus:    consts.MessageCreated,
			ServiceType:   consts.ServiceTypeApproval,
			BusinessId:    auditId,
		},
	}
	messageId, _ := s.MessageRepo.Create(messageReq)
	message, err := s.MessageRepo.Get(messageId)
	if err != nil {
		return
	}

	_, err = s.MessageService.SendMessageToMcs(message)

	return
}
