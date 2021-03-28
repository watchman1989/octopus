package redis

import (
	"fmt"
	"github.com/octopus/bootloader"
)

type RdsLoader struct {
	bootloader.BaseLoader
}


func (rl *RdsLoader) Setup(ctx *bootloader.LoadContext) {
	v := ctx.GetViper()
	host := v.GetString("redis.host")
	ctx.SetCtx("redis_client", "RC")

	fmt.Printf("redis loader init, host %s\n", host)
}


func (rl *RdsLoader) Stop(ctx *bootloader.LoadContext) {
	fmt.Printf("redis loader stop\n")
}


func (rl *RdsLoader) String() string {
	return "redis"
}