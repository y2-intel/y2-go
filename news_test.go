// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package y2_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/y2-intel/y2-go"
	"github.com/y2-intel/y2-go/internal/testutil"
	"github.com/y2-intel/y2-go/option"
)

func TestNewsListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := y2.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.News.List(context.TODO(), y2.NewsListParams{
		Limit:  y2.Int(1),
		Topics: y2.String("crypto,ai_agents,bitcoin"),
	})
	if err != nil {
		var apierr *y2.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNewsGetRecapsWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := y2.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.News.GetRecaps(context.TODO(), y2.NewsGetRecapsParams{
		Timeframe: y2.TimeframeEnum12h,
		Topics:    y2.String("topics"),
	})
	if err != nil {
		var apierr *y2.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNewsListFeeds(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := y2.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.News.ListFeeds(context.TODO())
	if err != nil {
		var apierr *y2.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
