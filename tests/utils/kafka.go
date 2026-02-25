package utils

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

var ErrorNotTheMessageIWasLookingFor = fmt.Errorf("not the message I was looking for")

type TestSubscriberConfig struct {
	Brokers []string
	Topic   string
	GroupID string
}

func NewTestSubscriberConfig(brokers []string, topic string, groupID string) TestSubscriberConfig {
	return TestSubscriberConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	}
}

func NewDefaultTestSubscriberConfig(topic string) TestSubscriberConfig {
	return TestSubscriberConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
		GroupID: fmt.Sprintf("test-group-%s-%d", topic, time.Now().UnixNano()),
	}
}

type TestSubscriber struct {
	reader  *kafka.Reader
	queue   chan kafka.Message
	errChan chan error
	stop    context.CancelFunc
}

func NewTestSubscriber(config TestSubscriberConfig) *TestSubscriber {
	return &TestSubscriber{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: config.Brokers,
			GroupID: config.GroupID,
			Topic:   config.Topic,
		}),
		queue:   make(chan kafka.Message),
		errChan: make(chan error, 1),
	}
}

func (ts *TestSubscriber) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	ts.stop = cancel

	go func() {
		for {
			m, err := ts.reader.ReadMessage(ctx)
			if err != nil {
				if !errors.Is(err, context.Canceled) {
					select {
					case ts.errChan <- err:
					default:
					}
				}
				return
			}

			select {
			case ts.queue <- m:
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (ts *TestSubscriber) WaitAndEvaluate(timeout time.Duration, evaluator func(kafka.Message) (any, error)) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for {
		select {
		case err := <-ts.errChan:
			return nil, fmt.Errorf("kafka error: %w", err)
		case <-ctx.Done():
			return nil, fmt.Errorf("timed out: %w", ctx.Err())
		case m := <-ts.queue:
			val, err := evaluator(m)
			if err != nil {
				if errors.Is(err, ErrorNotTheMessageIWasLookingFor) {
					continue
				}
				return nil, err
			}
			return val, nil
		}
	}
}

func (ts *TestSubscriber) Close() {
	if ts.stop != nil {
		ts.stop()
	}
	ts.reader.Close()
}

func CreateEmailSendRequestedMessageKey(firstName, lastName, email string) string {
	sanitizedFirstName := strings.ReplaceAll(firstName, ".", "")
	sanitizedFirstName = strings.ReplaceAll(sanitizedFirstName, "@", "")
	sanitizedLastName := strings.ReplaceAll(lastName, ".", "")
	sanitizedLastName = strings.ReplaceAll(sanitizedLastName, "@", "")

	recipient := fmt.Sprintf("%s %s <%s>", sanitizedFirstName, sanitizedLastName, email)

	hash := sha256.Sum256([]byte(recipient))
	id := hex.EncodeToString(hash[:])

	return id
}
