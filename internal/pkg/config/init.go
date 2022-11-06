package config

import (
	"bytes"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	myZap "github.com/aaronchen2k/deeptest/pkg/core/zap"
	_fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	_resUtils "github.com/aaronchen2k/deeptest/pkg/lib/res"
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
	if app == "agent" {
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
		yamlDefault, _ := _resUtils.ReadRes(configRes)

		if err := VIPER.ReadConfig(bytes.NewBuffer(yamlDefault)); err != nil {
			panic(fmt.Errorf("读取默认配置文件错误: %w ", err))
		}

		if err := VIPER.Unmarshal(&CONFIG); err != nil {
			panic(fmt.Errorf("同步配置文件错误: %w ", err))
		}

		if err := VIPER.WriteConfigAs(configFile); err != nil {
			panic(fmt.Errorf("写入配置文件错误: %w ", err))
		}
		return
	}

	// 存在配置文件，读取配置文件内容
	VIPER.SetConfigFile(configFile)
	err := VIPER.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置错误: %w ", err))
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

	if app == "server" {
		// 初始化Casbin配置
		casbinPath := consts.CasbinFileName
		fmt.Printf("Casbin配置文件为%s\n", casbinPath)

		if !dir.IsExist(casbinPath) {
			casbinRes := filepath.Join("res", consts.CasbinFileName)
			yamlDefault, _ := _resUtils.ReadRes(casbinRes)

			_, err = dir.WriteBytes(casbinPath, yamlDefault)
			if err != nil {
				panic(fmt.Errorf("初始化 casbin rbac_model.conf 文件错误: %w ", err))
			}
		}
	}
}
