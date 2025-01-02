package setting

import (
	"gin-blog/pkg/logging"
	"path/filepath"
	"runtime"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	PageNum   int
	JwtSecret string
)

func init() {
	var err error
	// 获取项目根目录
	_, filename, _, _ := runtime.Caller(0)
	rootPath := filepath.Join(filepath.Dir(filename), "../..")

	// 使用根目录构建配置文件路径
	configPath := filepath.Join(rootPath, "conf/app.ini")
	Cfg, err = ini.Load(configPath)
	if err != nil {
		logging.Fatal("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		logging.Fatal("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		logging.Fatal("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	PageNum = sec.Key("PAGE_NUM").MustInt(1)
}
