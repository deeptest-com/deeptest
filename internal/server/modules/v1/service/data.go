package service

import (
	"errors"
	logUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/log"
	"github.com/aaronchen2k/deeptest/internal/server/config"
	"github.com/aaronchen2k/deeptest/internal/server/core/cache"
	"github.com/aaronchen2k/deeptest/internal/server/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/source"
	"github.com/snowlyg/helper/str"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	ErrViperEmpty = errors.New("配置服务未初始化")
)

type DataService struct {
	DataRepo          *repo.DataRepo            `inject:""`
	UserRepo          *repo.UserRepo            `inject:""`
	UserSource        *source.UserSource        `inject:""`
	RoleSource        *source.RoleSource        `inject:""`
	PermSource        *source.PermSource        `inject:""`
	ProjectRoleSource *source.ProjectRoleSource `inject:""`
}

func NewDataService() *DataService {
	return &DataService{}
}

// writeConfig 写入配置文件
func (s *DataService) writeConfig(viper *viper.Viper, conf serverConfig.Config) error {
	cs := str.StructToMap(serverConfig.CONFIG)
	for k, v := range cs {
		viper.Set(k, v)
	}
	return viper.WriteConfig()
}

// 回滚配置
func (s *DataService) refreshConfig(viper *viper.Viper, conf serverConfig.Config) error {
	err := s.writeConfig(viper, conf)
	if err != nil {
		logUtils.Errorf("还原配置文件设置错误", zap.String("refreshConfig(consts.VIPER)", err.Error()))
		return err
	}
	return nil
}

// InitDB 创建数据库并初始化
func (s *DataService) InitDB(req serverDomain.DataReq) error {
	defaultConfig := serverConfig.CONFIG
	if serverConfig.VIPER == nil {
		logUtils.Errorf("初始化错误", zap.String("InitDB", ErrViperEmpty.Error()))
		return ErrViperEmpty
	}

	if serverConfig.CONFIG.System.CacheType == "redis" {
		serverConfig.CONFIG.Redis = serverConfig.Redis{
			DB:       serverConfig.CONFIG.Redis.DB,
			Addr:     serverConfig.CONFIG.Redis.Addr,
			Password: serverConfig.CONFIG.Redis.Password,
		}
		err := cache.Init() // redis缓存
		if err != nil {
			logUtils.Errorf("认证驱动初始化错误", zap.String("cache.Init() ", err.Error()))
			return err
		}
	}

	if serverConfig.CONFIG.System.DbType == "mysql" {
		if err := s.DataRepo.CreateMySqlDb(); err != nil {
			return err
		}
	}

	if err := s.writeConfig(serverConfig.VIPER, serverConfig.CONFIG); err != nil {
		logUtils.Errorf("更新配置文件错误", zap.String("writeConfig(consts.VIPER)", err.Error()))
	}

	if s.DataRepo.DB == nil {
		logUtils.Error("数据库初始化错误")
		s.refreshConfig(serverConfig.VIPER, defaultConfig)
		return errors.New("数据库初始化错误")
	}

	err := s.DataRepo.DB.AutoMigrate(model.Models...)
	if err != nil {
		logUtils.Errorf("迁移数据表错误", zap.String("错误:", err.Error()))
		s.refreshConfig(serverConfig.VIPER, defaultConfig)
		return err
	}

	if req.ClearData {
		err = s.initData(
			s.PermSource,
			s.RoleSource,
			s.ProjectRoleSource,
			s.UserSource,
		)
		if err != nil {
			logUtils.Errorf("填充数据错误", zap.String("错误:", err.Error()))
			s.refreshConfig(serverConfig.VIPER, defaultConfig)
			return err
		}
	}

	return nil
}

// initDB 初始化数据
func (s *DataService) initData(InitDBFunctions ...module.InitDBFunc) error {
	for _, v := range InitDBFunctions {
		err := v.Init()
		if err != nil {
			return err
		}
	}
	return nil
}
