package service

import (
	"bufio"
	"io"
	"os"
	"strings"
	"sync"
)

type config struct {
	// 配置文件路径
	ConfigPath string
	// 配置map
	Configs map[string]string
}

const DEV_PATH = "develop.env"
const PRO_PATH = "product.env"

var ConfigInstance *config
var muc sync.Mutex

func ShareConfigInstance(isProduct bool) *config {
	muc.Lock()
	defer muc.Unlock()

	if ConfigInstance == nil {
		configPath := PRO_PATH
		if isProduct {
			configPath = PRO_PATH
		} else {
			configPath = DEV_PATH
		}
		ConfigInstance = NewConfig(configPath).GetEnvConfig()
	}
	return ConfigInstance
}

func NewConfig(configPath string) *config {
	return &config{ConfigPath: configPath}
}

func (this *config) SetConfigPath(configPath string) {
	this.ConfigPath = configPath
}

func (this *config) GetConfigPath() string {
	return this.ConfigPath
}

func (this *config) GetEnvConfig() *config {
	config := initConfig(this.ConfigPath)
	if len(config) == 0 {
		panic("develop.env file can not be null")
	}
	this.Configs = config
	return this
}

func (this *config) GetConfigFromKey(key string) string {
	if len(this.Configs) == 0 {
		panic("no config initd")
	}
	if _, ok := this.Configs[key]; ok {
		return this.Configs[key]
	}
	return ""
}

func initConfig(path string) map[string]string {
	config := make(map[string]string)
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	return config
}
