package service

import (
	"errors"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	commService "github.com/aaronchen2k/deeptest/internal/pkg/service"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	_i118Utils "github.com/aaronchen2k/deeptest/pkg/lib/i118"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	_mailUtils "github.com/aaronchen2k/deeptest/pkg/lib/mail"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

var (
	ErrUserNameOrPassword = errors.New("用户名或密码错误")
)

type AccountService struct {
	UserRepo    *repo.UserRepo           `inject:""`
	LdapService *commService.LdapService `inject:""`
}

// Login 登录
func (s *AccountService) Login(tenantId consts.TenantId, req v1.LoginReq) (ret v1.LoginResp, err error) {
	var Id uint
	var userBase v1.UserBase
	var user model.SysUser
	user, _ = s.UserRepo.GetByUserName(tenantId, req.Username)
	if config.CONFIG.Ldap && req.Username != "admin" && !user.Type {
		userBase, err = s.LdapService.LdapUserInfo(req)
		if err != nil {
			return
		}
		Id, err = s.UserRepo.UpdateByLdapInfo(tenantId, userBase)
	} else {
		Id, err = s.UserLogin(tenantId, req)
	}

	if err != nil {
		return
	}

	claims := &multi.CustomClaims{
		ID:            strconv.FormatUint(uint64(Id), 10),
		Username:      req.Username,
		AuthorityId:   string(tenantId),
		AuthorityType: multi.AdminAuthority,
		LoginType:     multi.LoginTypeApp,
		AuthType:      multi.AuthPwd,
		CreationDate:  time.Now().Local().Unix(),
		ExpiresIn:     multi.RedisSessionTimeoutWeb.Milliseconds(),
	}

	ret.Token, _, err = multi.AuthDriver.GenerateToken(claims)
	if err != nil {
		return
	}

	return
}

func (s *AccountService) UserLogin(tenantId consts.TenantId, req v1.LoginReq) (userId uint, err error) {
	user, err := s.UserRepo.FindPasswordByUserName(tenantId, req.Username)
	if err != nil {
		user, err = s.UserRepo.FindPasswordByEmail(tenantId, req.Username)
		if err != nil {
			return
		}
	}
	userId = user.Id

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		logUtils.Errorf("用户名或密码错误", zap.String("密码:", req.Password), zap.String("hash:", user.Password), zap.String("bcrypt.CompareHashAndPassword()", err.Error()))
		err = ErrUserNameOrPassword
		return
	}
	return
}

func (s *AccountService) Register(tenantId consts.TenantId, req v1.RegisterReq) (err _domain.BizErr) {
	err = _domain.NoErr

	if req.Password != req.Confirm {
		err = _domain.ErrPasswordMustBeSame
		return
	}

	po, _ := s.UserRepo.FindByUserName(tenantId, req.Username)
	if po.Id > 0 {
		err = _domain.ErrUsernameExist
		return
	}

	user := model.SysUser{
		UserBase: v1.UserBase{
			Username: req.Username,
			Email:    req.Email,
		},
		Password: req.Password,
	}

	s.UserRepo.Register(tenantId, &user)

	if config.CONFIG.System.Name == "deeptest" {
		mp := map[string]string{
			"name": user.Name,
			"sys":  config.CONFIG.System.Name,
			"url":  config.CONFIG.System.Website,
		}
		_mailUtils.Send(user.Email, _i118Utils.Sprintf("register_success"), "register-success", mp)
	}

	return
}

func (s *AccountService) ForgotPassword(tenantId consts.TenantId, usernameOrPassword string) (err error) {
	user, err := s.UserRepo.GetByUsernameOrPassword(tenantId, usernameOrPassword)

	vcode, err := s.UserRepo.GenAndUpdateVcode(tenantId, user.ID)

	url := config.CONFIG.System.Website
	if !consts.IsRelease {
		url = consts.WebsiteDev
	}
	settings := map[string]string{
		"name":  user.Username,
		"url":   url,
		"vcode": vcode,
	}
	_mailUtils.Send(user.Email, _i118Utils.Sprintf("reset_password"), "reset-password", settings)

	return
}

func (s *AccountService) ResetPassword(tenantId consts.TenantId, req v1.ResetPasswordReq) (err error) {
	user, err := s.UserRepo.FindByUserNameAndVcode(tenantId, req.Username, req.Vcode)
	if err != nil {
		return
	}

	err = s.UserRepo.UpdatePassword(tenantId, req.Password, user.ID)
	if err != nil {
		return
	}

	err = s.UserRepo.ClearVcode(tenantId, user.ID)
	if err != nil {
		return
	}

	return
}
