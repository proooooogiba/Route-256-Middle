package producer

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/infra/kafka"
	"gitlab.ozon.dev/ipogiba/homework/loms/infra/kafka/producer"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
	"sync"
	"time"
)

// SyncProducer ...
type SyncProducer struct {
	mu       sync.RWMutex
	producer sarama.SyncProducer
}

func NewSyncProducer() (*SyncProducer, error) {
	config := kafka.Config{
		Brokers: []string{
			"localhost:9092",
		},
	}
	syncProducer, err := producer.NewSyncProducer(config,
		producer.WithIdempotent(),
		producer.WithRequiredAcks(sarama.WaitForAll),
		producer.WithMaxOpenRequests(1),
		producer.WithMaxRetries(5),
		producer.WithRetryBackoff(10*time.Millisecond),
	)
	return &SyncProducer{
		producer: syncProducer,
		mu:       sync.RWMutex{},
	}, err
}

func (sp *SyncProducer) SendMessages(ctx context.Context, msgs []model.ProducerMessage) error {
	saramaMsgs := make([]*sarama.ProducerMessage, len(msgs))
	for i, msg := range msgs {
		saramaMsgs[i] = &sarama.ProducerMessage{
			Topic: msg.Topic,
			Key:   sarama.StringEncoder(msg.Key),
			Value: sarama.StringEncoder(msg.Message),
		}
	}

	sp.mu.RLock()
	defer sp.mu.RUnlock()

	if ctxErr := ctx.Err(); ctxErr != nil {
		return ctxErr
	}
	err := sp.producer.SendMessages(saramaMsgs)
	return err
}

func (sp *SyncProducer) SendMessage(ctx context.Context, msg model.ProducerMessage) error {
	saramaMsg := &sarama.ProducerMessage{
		Topic: msg.Topic,
		Key:   sarama.StringEncoder(msg.Key),
		Value: sarama.StringEncoder(msg.Message),
	}

	sp.mu.RLock()
	defer sp.mu.RUnlock()

	if ctxErr := ctx.Err(); ctxErr != nil {
		return ctxErr
	}
	_, _, err := sp.producer.SendMessage(saramaMsg)
	return errors.Wrap(err, "failed to send message")
}
