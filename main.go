package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/huahuayu/address-generator/common/config"
	"github.com/huahuayu/address-generator/router"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

var (
	configPath string
	env        string
)

func init() {
	flag.StringVar(&configPath, "config", "", "config file path")
	flag.StringVar(&env, "env", "", "active env (dev, sit or prod)")
	flag.Parse()
}

func initTimezone() {
	os.Setenv("TZ", config.App.Server.TimezoneLoc)
}

func initLog() {
	logDir := config.App.Log.Path
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.MkdirAll(logDir, 0777)
		if err != nil {
			panic(fmt.Sprintf("create log dir err: %s\n", err))
		}
	}
	loc, _ := time.LoadLocation(config.App.Server.TimezoneLoc)
	t := time.Now().In(loc)
	file, err := os.OpenFile(fmt.Sprintf("%s%s%s.log", config.App.Log.Path, string(os.PathSeparator), fmt.Sprintf("%d-%02d-%02d",
		t.Year(), t.Month(), t.Day())), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0665)
	if err != nil {
		panic(fmt.Sprintf("open log file err: %s\n", err))
	}
	log.SetLevel(config.App.Log.Level)
	log.SetFormatter(&log.JSONFormatter{})
	out := io.MultiWriter(os.Stdout, file)
	log.SetOutput(out)
}

func main() {
	//config.Init(configPath, env)
	//initTimezone()
	//initLog()
	//db.Init()
	//redis.Init()

	gin.SetMode("debug")
	r := gin.Default()
	r.RedirectTrailingSlash = false

	router.Init(r)

	r.Run(":" + "8080")
}
