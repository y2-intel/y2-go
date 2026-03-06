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

func TestOsintSourceGetDataSourceHealth(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Osint.Sources.GetDataSourceHealth(context.TODO())
	if err != nil {
		var apierr *y2.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
