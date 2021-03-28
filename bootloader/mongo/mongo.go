package mongo

import (
	"fmt"
	"github.com/octopus/bootloader"
)

type MgoLoader struct {
	bootloader.BaseLoader
}


func (ml *MgoLoader) Setup(ctx *bootloader.LoadContext) {

	v := ctx.GetViper()
	host := v.GetString("mongo.host")

	ctx.SetCtx("mongo_client", "MC")

	fmt.Printf("Mongo loader init, host %s\n", host)
}


func (ml *MgoLoader) Stop(ctx *bootloader.LoadContext) {
	fmt.Printf("Mongo loader stop\n")
}

func (ml *MgoLoader) String() string {
	return "mongo"
}
