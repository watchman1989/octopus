package kafka

import (
	"fmt"
	"github.com/octopus/bootloader"
)

type KafLoader struct {
	bootloader.BaseLoader
}


func (kl *KafLoader) Setup(ctx *bootloader.LoadContext) {
	v := ctx.GetViper()
	cluster := v.GetString("kafka.cluster")
	topic := v.GetString("kafka.topic")

	ctx.SetCtx("kafka_client", "KC")

	fmt.Printf("Kafka loader init, cluster %s topic %s\n", cluster, topic)
}


func (kl *KafLoader) Stop(ctx *bootloader.LoadContext) {
	fmt.Printf("Kafka loader stop\n")
}

func (kl *KafLoader) String() string {
	return "kafka"
}