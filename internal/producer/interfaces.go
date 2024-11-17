package producer

import (
	"github.com/IBM/sarama"
)

// Producer интерфейс который должен реализовать продюсер брокер.
type Producer interface {
	SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error)
	Close() error
}
