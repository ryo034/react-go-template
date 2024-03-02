package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
)

type otpRateLimiter struct {
	rdb         *redis.Client
	limitConfig LimitConfig
}

type LimitConfig struct {
	Rate   int
	Burst  int
	Period time.Duration
}

func NewOtpRateLimitMiddleware(rdb *redis.Client, limitConfig LimitConfig) Middleware {
	return &otpRateLimiter{rdb, limitConfig}
}

// OtpRateLimitConfig Allow up to 5 requests per minute
func OtpRateLimitConfig() LimitConfig {
	r := redis_rate.PerMinute(5)
	return LimitConfig{r.Rate, r.Burst, r.Period}
}

type body struct {
	Email string `json:"email"`
}

const otpRateLimitKey = "otp-limit"

func (rl *otpRateLimiter) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// only limit otp endpoints
		lc := redis_rate.Limit{
			Rate:   rl.limitConfig.Rate,
			Burst:  rl.limitConfig.Burst,
			Period: rl.limitConfig.Period,
		}
		if r.URL.Path == "/otp/auth" || r.URL.Path == "/otp/verify" {
			b, err := io.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			defer func(Body io.ReadCloser) {
				if err = Body.Close(); err != nil {
					panic(err)
				}
			}(r.Body)

			var bd *body
			if err = json.Unmarshal(b, &bd); err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			if _, err = account.NewEmail(bd.Email); err != nil {
				h.ServeHTTP(w, r)
				return
			}

			rr := redis_rate.NewLimiter(rl.rdb)
			res, err := rr.Allow(r.Context(), fmt.Sprintf("%s:%s:%s", otpRateLimitKey, r.URL.Path, bd.Email), lc)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			if res.Allowed == 0 {
				http.Error(w, "The API is at capacity, try again later.", http.StatusTooManyRequests)
				return
			}
			r.Body = io.NopCloser(bytes.NewBuffer(b))
		}
		h.ServeHTTP(w, r)
	})
}
