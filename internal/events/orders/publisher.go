package orders

import (
	"context"
	"encoding/json"
	"first_go_project/internal/redis"
	"fmt"
)

const OrderCreatedChannel = "order.created"

func PublishOrderCreated(event OrderCreatedEvent) {
	data, _ := json.Marshal(event)
	err := redis.Rdb.Publish(context.Background(), OrderCreatedChannel, data).Err()
	if err != nil {
		fmt.Println("❌ Failed to publish order event:", err)
	}
}
