package keyvalue

import (
	"context"
	"github.com/go-faster/errors"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/ryo034/react-go-template/apps/system/api/util/test"
	"testing"
	"time"
)

func Test_redis_OK(t *testing.T) {
	ctx := context.Background()
	wantErr := false

	const key = "otp:test@example.com"
	const value = "123456"

	t.Run("OK", func(t *testing.T) {
		redisClient, err := test.SetupRedisClient(t, ctx)
		if err != nil {
			t.Errorf("SetupRedisClient() error = %v, wantErr %v", err, wantErr)
			return
		}

		c := NewRedisDriver(redisClient)

		err = c.Set(ctx, key, value, 120*time.Second)
		if (err != nil) != wantErr {
			t.Errorf("Set() error = %v, wantErr %v", err, wantErr)
			return
		}

		got, err := c.Get(ctx, key)
		if (err != nil) != wantErr {
			t.Errorf("Get() error = %v, wantErr %v", err, wantErr)
			return
		}
		if got != value {
			t.Errorf("Get() got = %v, want %v", got, value)
		}

		err = c.Delete(ctx, key)
		if (err != nil) != wantErr {
			t.Errorf("Delete() error = %v, wantErr %v", err, wantErr)
			return
		}

		got, err = c.Get(ctx, key)
		if !errors.Is(err, redis.Nil) {
			t.Errorf("Get() error = %v, wantErr %v", err, wantErr)
			return
		}
		if got != "" {
			t.Errorf("Get() got = %v, want %v", got, "")
		}
	})
}
