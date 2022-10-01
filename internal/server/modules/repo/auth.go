package repo

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthRepo struct {
	DB *gorm.DB `inject:""`
}

func NewAuthRepo() *AuthRepo {
	return &AuthRepo{}
}

func (r *AuthRepo) CreateToken(name, token, tokenType string, projectId int) (po model.Auth2Token, err error) {
	pos, _ := r.FindByToken(token)
	if len(pos) > 0 {
		err = errors.New("Token值已存在")
		return
	}

	pos, _ = r.FindByName(name)
	if len(pos) > 0 {
		r.RemoveToken(pos[0].ID)
	}

	po = model.Auth2Token{
		Name:      name,
		Token:     token,
		TokenType: tokenType,
		ProjectId: projectId}

	err = r.DB.Model(&po).Create(&po).Error
	if err != nil {
		logUtils.Errorf("add token error", zap.String("error:", err.Error()))
		err = fmt.Errorf("%d", _domain.ErrNameExist.Code)

		return
	}

	return
}

func (r *AuthRepo) ListOAuth2Token(projectId int) (pos []model.Auth2Token, err error) {
	err = r.DB.
		Where("project_id=?", projectId).
		Where("NOT deleted").
		Order("created_at DESC").
		Find(&pos).Error
	return

	return
}

func (r *AuthRepo) FindByToken(token string) (pos []model.Auth2Token, err error) {
	err = r.DB.
		Where("token=?", token).
		Where("NOT deleted").
		Find(&pos).Error
	return
}
func (r *AuthRepo) FindByName(name string) (pos []model.Auth2Token, err error) {
	err = r.DB.
		Where("name=?", name).
		Where("NOT deleted").
		Find(&pos).Error
	return
}

func (r *AuthRepo) RemoveToken(id uint) (err error) {
	err = r.DB.Model(&model.Auth2Token{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete token by id error", zap.String("error:", err.Error()))
		return
	}

	return
}
