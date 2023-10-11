package api

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
	"time"
)

type MessageHandler interface {
	HandleEvent(ev *kafka.Message)
}

type Consumer struct {
	topic   string
	c       *kafka.Consumer
	handler MessageHandler
}

func NewConsumer(topic string, handler MessageHandler) *Consumer {
	return &Consumer{topic: topic, handler: handler}
}

func (consumer *Consumer) Init(conf *kafka.ConfigMap) {
	c, err := kafka.NewConsumer(conf)

	if err != nil {
		logrus.Fatalf("Cannot create consumer: %s", err.Error())
	}

	err = c.Subscribe(consumer.topic, nil)
	if err != nil {
		logrus.Fatalf("Cannot subscribe to topic %s: %s", consumer.topic, err.Error())
	}

	consumer.c = c
}

func (consumer *Consumer) ReadMessages() {
	for {
		ev, err := consumer.c.ReadMessage(100 * time.Millisecond)

		if err != nil {
			continue
		}

		consumer.handler.HandleEvent(ev)
	}
}
