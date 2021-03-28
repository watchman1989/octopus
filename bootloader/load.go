package bootloader

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type LoadContext struct {
	mu sync.Mutex
	//config
	v *viper.Viper
	//context
	ctx map[string]interface{}
}

func NewLoadContext() *LoadContext {
	lc := &LoadContext{
		ctx: make(map[string]interface{}),
	}
	return lc
}

func (lc *LoadContext) SetViper(v *viper.Viper) {
	lc.v = v
}

func (lc *LoadContext) GetViper() *viper.Viper {
	return lc.v
}

func (lc *LoadContext) SetCtx(k string, v interface{}) {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	lc.ctx[k] = v
}

func (lc *LoadContext) GetCtx(k string) interface{} {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	return lc.ctx[k]
}


type Loader interface {
	Setup(*LoadContext)
	Start(*LoadContext)
	Stop(*LoadContext)
	String() string
}


type LoaderRegister struct {
	mu sync.Mutex
	Loaders []Loader
}

func newLoaderRegister() *LoaderRegister {
	sr := &LoaderRegister{Loaders: make([]Loader, 0)}
	return sr
}

func (sr *LoaderRegister) Register(Loader Loader) {
	sr.mu.Lock()
	sr.Loaders = append(sr.Loaders, Loader)
	sr.mu.Unlock()
	fmt.Printf("Register Loader: %s\n", Loader.String())
}


func (sr *LoaderRegister) GetLoaders() []Loader {
	Loaders := make([]Loader, 0)
	sr.mu.Lock()
	Loaders = append(Loaders, sr.Loaders...)
	sr.mu.Unlock()
	return Loaders
}


var loaderRegister = newLoaderRegister()

func Register(Loader Loader) {
	loaderRegister.Register(Loader)
}

func GetLoaders() []Loader {
	return loaderRegister.GetLoaders()
}

//base Loader

const (
	defaultLoaderString = "DefaultLoader"
)

type BaseLoader struct {}

func (bs *BaseLoader) Setup(lc *LoadContext) {}
func (bs *BaseLoader) Start(lc *LoadContext) {}
func (bs *BaseLoader) Stop(lc *LoadContext) {}
func (bs *BaseLoader) String() string {return defaultLoaderString}
