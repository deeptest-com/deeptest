package config

import (
	"bytes"
	"fmt"
	"github.com/aaronchen2k/deeptest"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	myZap "github.com/aaronchen2k/deeptest/pkg/core/zap"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	_fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	"github.com/go-redis/redis/v8"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/snowlyg/helper/dir"
	"github.com/spf13/viper"
)

var (
	CONFIG     Config
	VIPER      *viper.Viper
	CACHE      redis.UniversalClient
	PermRoutes []map[string]string
)

func Init(app string) {
	consts.IsRelease = _commUtils.IsRelease()

	if app == "server" {
		// 初始化Casbin配置
		casbinPath := consts.CasbinFileName
		fmt.Sprintf("casbin conf file is %s", casbinPath)

		if !dir.IsExist(casbinPath) {
			casbinRes := filepath.Join("res", consts.CasbinFileName)
			yamlDefault, err := deeptest.ReadResData(casbinRes)
			if err != nil {
				panic(fmt.Errorf("failed to read casbin rbac_model.conf from res: %s", err.Error()))
			}

			err = _fileUtils.WriteFile(casbinPath, string(yamlDefault))
			if err != nil {
				panic(fmt.Errorf("failed to write casbin rbac_model.conf 文件错误: %s", err.Error()))
			}
		}
	} else if app == "agent" {
		home, _ := _fileUtils.GetUserHome()
		consts.HomeDir = filepath.Join(home, consts.App)
		consts.TmpDir = filepath.Join(consts.HomeDir, consts.FolderTmp)

		_fileUtils.MkDirIfNeeded(consts.TmpDir)
	}

	v := viper.New()
	VIPER = v
	VIPER.SetConfigType("yaml")

	configFile := consts.ConfigFileName
	fmt.Printf("配置文件路径为%s\n", configFile)

	if !dir.IsExist(configFile) { // 没有配置文件，写入默认配置
		configRes := filepath.Join("res", consts.ConfigFileName)
		yamlDefault, _ := deeptest.ReadResData(configRes)

		if err := VIPER.ReadConfig(bytes.NewBuffer(yamlDefault)); err != nil {
			panic(fmt.Errorf("读取默认配置文件错误: %w ", err))
		}
		if err := VIPER.Unmarshal(&CONFIG); err != nil {
			panic(fmt.Errorf("同步配置文件错误: %w ", err))
		}
		if err := VIPER.WriteConfigAs(configFile); err != nil {
			panic(fmt.Errorf("写入配置文件错误: %w ", err))
		}
	} else {
		VIPER.SetConfigFile(configFile)
		err := VIPER.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("读取配置错误: %w ", err))
		}
	}

	// 监控配置文件变化
	VIPER.WatchConfig()
	VIPER.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置发生变化:", e.Name)
		if err := VIPER.Unmarshal(&CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&CONFIG); err != nil {
		fmt.Println(err)
	}
	myZap.ZapInst = CONFIG.Zap
}
