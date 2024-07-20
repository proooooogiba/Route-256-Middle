package order

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
)

func (c *Order) sendOrderEvent(ctx context.Context, order *model.Order, status model.Status) error {
	event, err := json.Marshal(&model.OrderEvent{
		ID:              order.ID,
		EventType:       status,
		OperationMoment: time.Now(),
		ExtraInfo:       "",
	})
	if err != nil {
		return errors.Wrap(err, "json.Marshal")
	}

	err = c.producer.SendMessage(ctx, model.ProducerMessage{
		Topic:   c.topic,
		Key:     strconv.FormatInt(order.ID, 10),
		Message: string(event),
	})

	return errors.Wrap(err, "producer.SendMessage")
}
