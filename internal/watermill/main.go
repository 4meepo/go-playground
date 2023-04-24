package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/redis/go-redis/v9"
)

func main() {
	subClient := redis.NewClient(&redis.Options{
		Addr: "ecs:6379",
		DB:   0,
	})
	{
		consumer(subClient, "test_consumer_group")
		consumer(subClient, "test_consumer_group2")
	}

	pubClient := redis.NewClient(&redis.Options{
		Addr: "ecs:6379",
		DB:   0,
	})
	publisher, err := redisstream.NewPublisher(
		redisstream.PublisherConfig{
			Client:     pubClient,
			Marshaller: redisstream.DefaultMarshallerUnmarshaller{},
			Maxlens: map[string]int64{
				"example.topic": 50,
			},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	publishMessages(publisher)
}

func consumer(subClient *redis.Client, group string) {
	subscriber, err := redisstream.NewSubscriber(
		redisstream.SubscriberConfig{
			Client:        subClient,
			Unmarshaller:  redisstream.DefaultMarshallerUnmarshaller{},
			ConsumerGroup: group,
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	{
		messages, err := subscriber.Subscribe(context.Background(), "example.topic")
		if err != nil {
			panic(err)
		}

		go process("c1", messages)
	}
	{
		messages, err := subscriber.Subscribe(context.Background(), "example.topic")
		if err != nil {
			panic(err)
		}

		go process("c2", messages)
	}
}

func publishMessages(publisher message.Publisher) {
	i := 0
	for {
		msg := message.NewMessage(watermill.NewUUID(), []byte(fmt.Sprintf("Hello, world ! %d", i)))
		i++
		if err := publisher.Publish("example.topic", msg); err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}

func process(id string, messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("%s received message: %s, payload: %s", id, msg.UUID, string(msg.Payload))

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}
