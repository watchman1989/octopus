package bootloader

import (
	"fmt"
	"github.com/spf13/viper"
)

type Boot struct {
	conf       *viper.Viper
	loadContext *LoadContext
}

func NewBoot(conf *viper.Viper) *Boot {
	b := &Boot{
		conf:       conf,
		loadContext: NewLoadContext(),
	}
	b.loadContext.SetViper(conf)
	return b
}

func (b *Boot) Init() {
	fmt.Printf("init loaders ...\n")
	for _, v := range GetLoaders() {
		fmt.Printf(">>>>init loader %s\n", v.String())
		v.Setup(b.loadContext)
	}
}

func (b *Boot) Start() {
}

func (b *Boot) Stop() {
	fmt.Printf("stop loaders ...\n")
	for _, v := range GetLoaders() {
		fmt.Printf("stop loader %s\n", v.String())
		v.Stop(b.loadContext)
	}
}
