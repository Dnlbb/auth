package config

import (
	"fmt"
	"os"
	"strings"
)

const (
	kafkaAddress = "BROKER_ADDRESS"
	topic        = "TOPIC"
)

// KafkaConf конфиг для кафки.
type KafkaConf interface {
	Addresses() []string
	Topic() string
}

type kafkaConfImpl struct {
	addresses []string
	topic     string
}

// NewKafkaConfImpl конструктор для конфига.
func NewKafkaConfImpl() (KafkaConf, error) {
	address := os.Getenv(kafkaAddress)
	if len(address) == 0 {
		return nil, fmt.Errorf("environment variable %s not defined", kafkaAddress)
	}

	SliceAddress := strings.Split(address, ",")

	topic := os.Getenv(topic)
	if len(topic) == 0 {
		return nil, fmt.Errorf("environment variable %s not defined", topic)
	}

	return kafkaConfImpl{
		addresses: SliceAddress,
		topic:     topic,
	}, nil
}

func (k kafkaConfImpl) Addresses() []string {
	return k.addresses
}

func (k kafkaConfImpl) Topic() string {
	return k.topic
}
