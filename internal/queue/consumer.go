package queue

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/dexguitar/chatapp/configs"
)

type KafkaConsumer struct {
	sarama.ConsumerGroup
}

func NewKafkaConsumer(appConfig *configs.Config) (*KafkaConsumer, error) {
	op := "queue.NewKafkaConsumer"

	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumerGroup(appConfig.KafkaBrokers, appConfig.ConsumerGroup, config)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &KafkaConsumer{
		ConsumerGroup: consumer,
	}, nil
}

func (c *KafkaConsumer) ConsumeMessages(ctx context.Context, h Hub) {
	for {
		err := c.ConsumerGroup.Consume(ctx, []string{topic}, &ConsumerHandler{Hub: h})
		if err != nil {
			return
		}
	}
}
