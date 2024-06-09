package queue

import (
	"context"
	"fmt"

	"github.com/dexguitar/chatapp/internal/model"
	"github.com/snaffi/errors"
)

type Queue struct {
	Producer *KafkaProducer
	Consumer *KafkaConsumer
}

func New(p *KafkaProducer, c *KafkaConsumer) (*Queue, error) {
	if p == nil || c == nil {
		return nil, errors.New("producer or consumer cannot be nil")
	}

	return &Queue{
		Producer: p,
		Consumer: c,
	}, nil
}

func (q *Queue) WriteMessage(ctx context.Context, message *model.Message) error {
	op := "Queue.WriteMessage"

	err := q.Producer.ProduceMessage(ctx, message)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
