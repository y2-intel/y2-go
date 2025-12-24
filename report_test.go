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

func TestReportGet(t *testing.T) {
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
	_, err := client.Reports.Get(context.TODO(), "k57abc123def456")
	if err != nil {
		var apierr *y2.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReportListWithOptionalParams(t *testing.T) {
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
	_, err := client.Reports.List(context.TODO(), y2.ReportListParams{
		Limit:     y2.Int(1),
		ProfileID: y2.String("k57abc123def456"),
	})
	if err != nil {
		var apierr *y2.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReportGetAudioWithOptionalParams(t *testing.T) {
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
	_, err := client.Reports.GetAudio(
		context.TODO(),
		"reportId",
		y2.ReportGetAudioParams{
			Redirect: y2.Bool(true),
		},
	)
	if err != nil {
		var apierr *y2.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
