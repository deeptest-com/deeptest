package service

import (
	"errors"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"

	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

var (
	ErrUserNameOrPassword = errors.New("用户名或密码错误")
)

type AccountService struct {
	UserRepo *repo.UserRepo `inject:""`
}

func NewAuthService() *AccountService {
	return &AccountService{}
}

// Login 登录
func (s *AccountService) Login(req v1.LoginReq) (token string, err error) {
	user, err := s.UserRepo.FindPasswordByUserName(req.Username)
	if err != nil {
		user, err = s.UserRepo.FindPasswordByEmail(req.Username)

		if err != nil {
			return
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		logUtils.Errorf("用户名或密码错误", zap.String("密码:", req.Password), zap.String("hash:", user.Password), zap.String("bcrypt.CompareHashAndPassword()", err.Error()))
		err = ErrUserNameOrPassword
		return
	}

	claims := &multi.CustomClaims{
		ID:            strconv.FormatUint(uint64(user.Id), 10),
		Username:      req.Username,
		AuthorityId:   "",
		AuthorityType: multi.AdminAuthority,
		LoginType:     multi.LoginTypeWeb,
		AuthType:      multi.AuthPwd,
		CreationDate:  time.Now().Local().Unix(),
		ExpiresIn:     multi.RedisSessionTimeoutWeb.Milliseconds(),
	}
	token, _, err = multi.AuthDriver.GenerateToken(claims)
	if err != nil {
		return "", err
	}

	return
}

func (s *AccountService) Register(req v1.RegisterReq) (err _domain.BizErr) {
	err = _domain.NoErr

	if req.Password != req.Confirm {
		err = _domain.ErrPasswordMustBeSame
		return
	}

	po, _ := s.UserRepo.FindByUserName(req.Username)
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

	s.UserRepo.Register(&user)

	return
}
