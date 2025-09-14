package orders

import (
	"context"
	"encoding/json"
	"first_go_project/internal/redis"
	"fmt"
)

func ListenOrderCreated() {
	sub := redis.Rdb.Subscribe(context.Background(), OrderCreatedChannel)
	ch := sub.Channel()

	go func() {
		for msg := range ch {
			var e OrderCreatedEvent
			if err := json.Unmarshal([]byte(msg.Payload), &e); err != nil {
				fmt.Println("❌ Failed to unmarshal order event:", err)
				continue
			}
			fmt.Printf("[EVENT] Order created: %+v\n", e)

			// здесь можно вызвать почтовую рассылку, нотификации, воркер и т.д.
		}
	}()
}
