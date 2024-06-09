package queue

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/dexguitar/chatapp/configs"
	"github.com/dexguitar/chatapp/internal/model"
)

type KafkaProducer struct {
	sarama.SyncProducer
}

func NewKafkaProducer(appConfig *configs.Config) (*KafkaProducer, error) {
	op := "queue.NewKafkaProducer"

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(appConfig.KafkaBrokers, config)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &KafkaProducer{
		SyncProducer: producer,
	}, nil
}

func (kw *KafkaProducer) ProduceMessage(ctx context.Context, msg *model.Message) error {
	op := "KafkaProducer.ProduceMessage"

	_, _, err := kw.SyncProducer.SendMessage(
		&sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(fmt.Sprintf("`%s` says `%s` to `%s`", msg.Username, msg.Value, msg.Receiver)),
			Headers: []sarama.RecordHeader{
				{Key: []byte("username"), Value: []byte(msg.Username)},
				{Key: []byte("value"), Value: []byte(msg.Value)},
				{Key: []byte("receiver"), Value: []byte(msg.Receiver)},
			},
		},
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
