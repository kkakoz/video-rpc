package producer

import (
	"github.com/Shopify/sarama"
	"github.com/kkakoz/pkg/kafkax"
	"github.com/spf13/viper"
)

var producer sarama.SyncProducer

func Init(viper *viper.Viper) (err error) {
	producer, err = kafkax.NewSyncProducer(viper)
	return err
}

func Send(value *sarama.ProducerMessage) error {
	//data, err := json.Marshal(event)
	//if err != nil {
	//	return err
	//}
	//
	//msg := &sarama.ProducerMessage{
	//	Topic: "video-web-event",
	//	Value: sarama.StringEncoder(data),
	//	//Key:   sarama.StringEncoder(strconv.Itoa(11)),
	//}

	_, _, err := producer.SendMessage(value)
	return err
}
