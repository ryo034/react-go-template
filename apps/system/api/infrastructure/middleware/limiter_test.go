//go:build testcontainers

package middleware

import (
	"context"
	"fmt"
	"github.com/go-redis/redis_rate/v10"
	"github.com/ryo034/react-go-template/apps/system/api/util/test"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_OrpRateLimiterMiddleware_Burst(t *testing.T) {
	// rate limit: 1 request per second
	ctx := context.Background()
	rc, err := test.SetupRedisClient(ctx)
	if err != nil {
		t.Fatalf("Failed to setup redis client: %v", err)
		return
	}

	pm := redis_rate.PerMinute(5)
	conf := LimitConfig{
		Rate:   pm.Rate,
		Burst:  pm.Burst,
		Period: pm.Period,
	}
	middleware := NewOtpRateLimitMiddleware(rc, conf)

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler := middleware.Handler(testHandler)

	server := httptest.NewServer(handler)
	defer server.Close()

	// Confirm that multiple users can access without affecting each other's rate limits
	body1 := `{"email":"test+1@example.com"}`
	body2 := `{"email":"test+2@example.com"}`

	// Under the limit
	authEp := "/otp/auth"
	for i := 0; i < 5; i++ {
		resp, err := http.Post(fmt.Sprintf("%s%s", server.URL, authEp), "application/json", strings.NewReader(body1))
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected OK, got %v", resp.Status)
		}
	}
	// Over the limit
	resp, err := http.Post(fmt.Sprintf("%s%s", server.URL, authEp), "application/json", strings.NewReader(body1))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	if resp.StatusCode != http.StatusTooManyRequests {
		t.Errorf("Expected TooManyRequests, got %v", resp.Status)
	}
	// 別のユーザーがアクセスしても、互いのレート制限に影響を与えないことを確認する
	resp, err = http.Post(fmt.Sprintf("%s%s", server.URL, authEp), "application/json", strings.NewReader(body2))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	if resp.StatusCode == http.StatusTooManyRequests {
		t.Errorf("Expected TooManyRequests, got %v", resp.Status)
	}

	// ===============
	// Verify Endpoint
	// ===============

	// Under the limit
	verifyEp := "/otp/verify"
	for i := 0; i < 5; i++ {
		resp, err = http.Post(fmt.Sprintf("%s%s", server.URL, verifyEp), "application/json", strings.NewReader(body1))
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected OK, got %v", resp.Status)
		}
	}
	// Over the limit
	resp, err = http.Post(fmt.Sprintf("%s%s", server.URL, verifyEp), "application/json", strings.NewReader(body1))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	if resp.StatusCode != http.StatusTooManyRequests {
		t.Errorf("Expected TooManyRequests, got %v", resp.Status)
	}
	// 別のユーザーがアクセスしても、互いのレート制限に影響を与えないことを確認する
	resp, err = http.Post(fmt.Sprintf("%s%s", server.URL, verifyEp), "application/json", strings.NewReader(body2))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	if resp.StatusCode == http.StatusTooManyRequests {
		t.Errorf("Expected TooManyRequests, got %v", resp.Status)
	}
}
