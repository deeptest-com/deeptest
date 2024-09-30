package repo

import (
	"errors"
	"fmt"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/pkg/domain"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *AuthRepo) CreateToken(tenantId consts.TenantId, name, token, tokenType string, projectId int) (po model.Auth2Token, err error) {
	pos, _ := r.FindByToken(tenantId, token)
	if len(pos) > 0 {
		err = errors.New("Token值已存在")
		return
	}

	pos, _ = r.FindByName(tenantId, name)
	if len(pos) > 0 {
		r.RemoveToken(tenantId, pos[0].ID)
	}

	po = model.Auth2Token{
		Name:      name,
		Token:     token,
		TokenType: tokenType,
		ProjectId: projectId}

	err = r.GetDB(tenantId).Model(&po).Create(&po).Error
	if err != nil {
		logUtils.Errorf("add token error", zap.String("error:", err.Error()))
		err = fmt.Errorf("%d", _domain.ErrNameExist.Code)

		return
	}

	return
}

func (r *AuthRepo) ListOAuth2Token(tenantId consts.TenantId, projectId int) (pos []model.Auth2Token, err error) {
	err = r.GetDB(tenantId).
		Where("project_id=?", projectId).
		Where("NOT deleted").
		Order("created_at DESC").
		Find(&pos).Error
	return

	return
}

func (r *AuthRepo) FindByToken(tenantId consts.TenantId, token string) (pos []model.Auth2Token, err error) {
	err = r.GetDB(tenantId).
		Where("token=?", token).
		Where("NOT deleted").
		Find(&pos).Error
	return
}
func (r *AuthRepo) FindByName(tenantId consts.TenantId, name string) (pos []model.Auth2Token, err error) {
	err = r.GetDB(tenantId).
		Where("name=?", name).
		Where("NOT deleted").
		Find(&pos).Error
	return
}

func (r *AuthRepo) RemoveToken(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.Auth2Token{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete token by id error", zap.String("error:", err.Error()))
		return
	}

	return
}
