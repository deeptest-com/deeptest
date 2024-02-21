package config

import (
	"bytes"
	"fmt"
	"github.com/aaronchen2k/deeptest"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	myZap "github.com/aaronchen2k/deeptest/pkg/core/zap"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	_fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"github.com/snowlyg/helper/dir"
	"github.com/spf13/viper"
	"path/filepath"
)

var (
	CONFIG     Config
	VIPER      *viper.Viper
	CACHE      redis.UniversalClient
	PermRoutes []map[string]string
)

func Init() {
	consts.IsRelease = _commUtils.IsRelease()

	v := viper.New()
	VIPER = v
	VIPER.SetConfigType("yaml")

	// agent
	if consts.RunFrom == consts.FromAgent {
		agentConfigPath := filepath.Join(consts.WorkDir, consts.AgentConfigFileName)

		if !dir.IsExist(agentConfigPath) { // 没有配置文件，写入默认配置
			configRes := filepath.Join("res", consts.AgentConfigFileName)
			yamlDefault, _ := deeptest.ReadResData(configRes)

			if err := VIPER.ReadConfig(bytes.NewBuffer(yamlDefault)); err != nil {
				panic(fmt.Errorf("读取默认配置文件错误: %w ", err))
			}

			if err := VIPER.WriteConfigAs(agentConfigPath); err != nil {
				panic(fmt.Errorf("写入配置文件错误: %w ", err))
			}
		}

		VIPER.SetConfigFile(agentConfigPath)
		err := VIPER.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("读取配置错误: %w ", err))
		}

		if consts.Port > 0 {
			CONFIG.System.AgentAddress = fmt.Sprintf("0.0.0.0:%d", consts.Port)
		}

		myZap.ZapInst = CONFIG.Zap

		if CONFIG.System.Name == "" {
			CONFIG.System.Name = consts.App
		}
		getWorkdir(CONFIG.System.Name)

		return
	}

	// server
	// 初始化rbac_model.conf配置
	casbinPath := filepath.Join(consts.WorkDir, consts.CasbinFileName)
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

	// 初始化server.yaml配置
	serverConfigPath := filepath.Join(consts.WorkDir, consts.ServerConfigFileName)
	if !dir.IsExist(serverConfigPath) { // 没有配置文件，写入默认配置
		configRes := filepath.Join("res", consts.ServerConfigFileName)
		yamlDefault, _ := deeptest.ReadResData(configRes)

		if err := VIPER.ReadConfig(bytes.NewBuffer(yamlDefault)); err != nil {
			panic(fmt.Errorf("读取默认配置文件错误: %w ", err))
		}

		if err := VIPER.WriteConfigAs(serverConfigPath); err != nil {
			panic(fmt.Errorf("写入配置文件错误: %w ", err))
		}
	}

	VIPER.SetConfigFile(serverConfigPath)
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

	if err := VIPER.Unmarshal(&CONFIG); err != nil {
		fmt.Println(err)
	}

	getWorkdir(CONFIG.System.Name)

	CONFIG.System.SysEnv = _commUtils.GetEnvVar("SysEnv", CONFIG.System.SysEnv)
	myZap.ZapInst = CONFIG.Zap
}

func getWorkdir(appName string) {
	home, _ := _fileUtils.GetUserHome()
	consts.WorkDir = filepath.Join(home, appName)
	consts.TmpDir = filepath.Join(consts.WorkDir, consts.FolderTmp)

	_fileUtils.MkDirIfNeeded(consts.TmpDir)
}
