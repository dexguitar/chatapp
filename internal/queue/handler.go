package queue

import (
	"github.com/IBM/sarama"
)

type ConsumerHandler struct {
	Hub Hub
}

// Setup is run at the beginning of a new session
func (h *ConsumerHandler) Setup(sarama.ConsumerGroupSession) error { return nil }

// Cleanup is run at the end of a session
func (h *ConsumerHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (h *ConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		if message.Value == nil {
			continue
		}
		h.Hub.Broadcast(message)
	}
	return nil
}
