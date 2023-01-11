package async

import (
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/kkakoz/pkg/app/kafkas"
	"github.com/kkakoz/pkg/logger"
	"github.com/kkakoz/video-rpc/pb/common"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type EventConsumer struct {
	*kafkas.KafkaConsumer
}

func NewEventConsumer(viper *viper.Viper) (*EventConsumer, error) {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true

	consumer, err := kafkas.NewConsumer(viper, consumerHandler, config)
	return &EventConsumer{KafkaConsumer: consumer}, err
}

func consumerHandler(message *sarama.ConsumerMessage) error {

	event := &common.Event{}
	err := proto.Unmarshal(message.Value, event)
	if err != nil {
		return err
	}

	logger.Debug("consumer event:" + string(message.Value))

	eventHandler := handlerMap[event.EventType]
	if eventHandler == nil {
		return nil
	}
	err = eventHandler(event)
	if err != nil {
		logger.Error("consumer event err", zap.Error(err), zap.String("event", string(message.Value)))
	}
	return nil
}
