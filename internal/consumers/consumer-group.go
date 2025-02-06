package consumers

import (
	"context"
	"fmt"
	"log/slog"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
)

type ConsumerGroup struct {
	consumer sarama.ConsumerGroupHandler
	group    string
	brokers  []string
	topic    string
	logger   *slog.Logger
}

func NewConsumerGroup(
	consumer sarama.ConsumerGroupHandler,
	group string,
	topic string,
	brokers []string,
	logger *slog.Logger,
) *ConsumerGroup {
	return &ConsumerGroup{
		consumer: consumer,
		group:    group,
		topic:    topic,
		brokers:  brokers,
		logger:   logger,
	}
}

func newSaramaConsumerGroup(cg *ConsumerGroup) (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	client, err := sarama.NewConsumerGroup(cg.brokers, cg.group, config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (cg *ConsumerGroup) Run(ctx context.Context) error {

	//TODO: вытащить конфиг во внешку, что бы было универсально для всего//проверить с Артуром

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	client, err := newSaramaConsumerGroup(cg)
	if err != nil {
		return fmt.Errorf("error creating consumer group client: %w", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, []string{cg.topic}, cg.consumer); err != nil { //
				cg.logger.Error("error consuming messages",
					slog.String("err", err.Error()))
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()
	cg.logger.Info("Consumer up and ready")

	//signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	interruptCtx, interruptCancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer interruptCancel()

	<-interruptCtx.Done()
	cg.logger.Info("recived interrupt signal", slog.String("err", interruptCtx.Err().Error()))

	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		return fmt.Errorf("error closing consumer group client: %w", err)
	}
	return nil

}
