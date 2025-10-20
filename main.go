package main

import (
	"fmt"
	"time"
	"utils/Log"
)

func main() {
	var config Log.RotateConfig
	config.FilePath = "./logger/"
	config.AppName = "app"
	cfg := Log.Init(&config, Log.WithMaxBackups(1000), Log.WithLevel("debug"), Log.WithStdout(true), Log.WithRotateTime(5), Log.WithRotateSize(1000))
	//cfgJson, _ := json.Marshal(cfg)
	logger := Log.New(cfg)
	fmt.Println("+++++++++++++++++")
	for i := 0; i < 50; i++ {
		logger.Info(fmt.Sprintf("%d === WWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWW", i))
		time.Sleep(time.Duration(1) * time.Second)
	}
	logger.Info("Warn msg")
	//logger.Info("Warn msg", Log.String("val", "string"))
	fmt.Println("+++++++++++++++++")

}
