package main

import (
	"flag"
	"fmt"
	"github.com/octopus/bootloader"
	"github.com/octopus/bootloader/kafka"
	"github.com/octopus/bootloader/mongo"
	"github.com/octopus/bootloader/redis"
	"github.com/spf13/viper"
	"time"
)


const (
	defaultConfigFile = "./conf/octopus.yaml"
)

var (
	v *viper.Viper
)


func init() {
	bootloader.Register(&redis.RdsLoader{})
	bootloader.Register(&mongo.MgoLoader{})
	bootloader.Register(&kafka.KafLoader{})

}

func main() {
	//get config path
	confPath := flag.String("c", defaultConfigFile, "ConfigPath")
	flag.Parse()
	//init config
	v = viper.New()
	v.SetConfigFile(*confPath)
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("read config file error: %s\n", err.Error())
		return
	}
	//new boot loader
	b := bootloader.NewBoot(v)
	b.Init()
	b.Start()

	time.Sleep(5 * time.Second)
	b.Stop()
}

