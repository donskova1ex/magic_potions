package consumers

import (
	"context"
	"log"
	"time"

	"github.com/IBM/sarama"
)

// TODO: проверить и задать вопросы
type Saver interface {
	Save(ctx context.Context, key []byte, body []byte, timeStamp time.Time) error
}

type RecipesConsumer struct {
	saver Saver
}

func NewRecipesConsumer(saver Saver) *RecipesConsumer {
	return &RecipesConsumer{saver}
}
func (c *RecipesConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}
func (c *RecipesConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// TODO: прокинуть свой логгер
func (c *RecipesConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Println("Consumer kafka messages chan closed")
				return nil
			}
			if err := c.saver.Save(session.Context(), message.Key, message.Value, message.Timestamp); err != nil {
				return err
			}
			session.MarkMessage(message, "read")
			session.Commit()
		case <-session.Context().Done():
			return nil
		}
	}
}
