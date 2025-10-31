package main

import (
	"fmt"
	"os"
	"reflect"
	//"utils/Log"
)

func main() {
	tty, err := os.OpenFile("./LogTest/Rotate.go_bak", os.O_RDWR, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		//aaa := reflect.TypeOf(tty)
		switch fmt.Sprintln(reflect.TypeOf(tty)) {
		case "*os.File":
			fmt.Println("11111111111")
		}
		//fmt.Sprintln(reflect.TypeOf(tty))
		fmt.Println("hello world")
	}
	//var config Log.RotateConfig
	//fmt.Println(config)
	//config.FilePath = "./logger/"
	//config.AppName = "app"
	//cfg := Log.Init(&config, Log.WithMaxBackups(5), Log.WithRotateSize(1000), Log.WithLocalTime(true), Log.WithLevel("debug"), Log.WithStdout(true), Log.WithRotateTime(5))
	//
	////多种文档输出
	//tee := []Log.TeeOption{
	//	{
	//		Out:              os.Stdout,
	//		LevelEnablerFunc: func(level Log.Level) bool { return level < Log.WarnLevel },
	//	}, {
	//		Out:              os.Stdout,
	//		LevelEnablerFunc: func(level Log.Level) bool { return level < Log.InfoLevel },
	//	},
	//}
	//Log.NewTee(cfg)
	//
	////cfgJson, _ := json.Marshal(cfg)
	//logger := Log.New(cfg)
	//fmt.Println("+++++++++++++++++")
	//for i := 0; i < 100000; i++ {
	//	logger.Info(fmt.Sprintf("%d === WWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWW", i))
	//	time.Sleep(time.Duration(1) * time.Millisecond)
	//}
	//logger.Info("Warn msg")
	//time.Sleep(2 * time.Second)
	////logger.Info("Warn msg", Log.String("val", "string"))
	//fmt.Println("+++++++++++++++++")

}
