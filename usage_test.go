// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package y2_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/y2-go"
	"github.com/stainless-sdks/y2-go/internal/testutil"
	"github.com/stainless-sdks/y2-go/option"
)

func TestUsage(t *testing.T) {
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
	t.Skip("Prism tests are disabled")
	reports, err := client.Reports.List(context.TODO(), y2.ReportListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", reports.Data)
}
