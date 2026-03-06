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

func TestOsintGetConflictIndicatorsWithOptionalParams(t *testing.T) {
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
	_, err := client.Osint.GetConflictIndicators(context.TODO(), y2.OsintGetConflictIndicatorsParams{
		Category: y2.OsintGetConflictIndicatorsParamsCategorySeismic,
		Limit:    y2.Int(1),
		Region:   y2.OsintGetConflictIndicatorsParamsRegionMena,
	})
	if err != nil {
		var apierr *y2.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOsintGetGpsJammingZonesWithOptionalParams(t *testing.T) {
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
	_, err := client.Osint.GetGpsJammingZones(context.TODO(), y2.OsintGetGpsJammingZonesParams{
		Limit:    y2.Int(1),
		Severity: y2.OsintGetGpsJammingZonesParamsSeverityLow,
	})
	if err != nil {
		var apierr *y2.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOsintGetMilitaryPostureWithOptionalParams(t *testing.T) {
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
	_, err := client.Osint.GetMilitaryPosture(context.TODO(), y2.OsintGetMilitaryPostureParams{
		Limit: y2.Int(1),
	})
	if err != nil {
		var apierr *y2.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOsintListAircraftWithOptionalParams(t *testing.T) {
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
	_, err := client.Osint.ListAircraft(context.TODO(), y2.OsintListAircraftParams{
		Limit:   y2.Int(1),
		Theater: y2.String("theater"),
	})
	if err != nil {
		var apierr *y2.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOsintListEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.Osint.ListEvents(context.TODO(), y2.OsintListEventsParams{
		Category: y2.OsintListEventsParamsCategorySeismic,
		Limit:    y2.Int(1),
		Severity: y2.OsintListEventsParamsSeverityLow,
	})
	if err != nil {
		var apierr *y2.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOsintListVesselsWithOptionalParams(t *testing.T) {
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
	_, err := client.Osint.ListVessels(context.TODO(), y2.OsintListVesselsParams{
		Limit:  y2.Int(1),
		Region: y2.String("region"),
	})
	if err != nil {
		var apierr *y2.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOsintMapEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.Osint.MapEvents(context.TODO(), y2.OsintMapEventsParams{
		Limit:  y2.Int(1),
		Region: y2.OsintMapEventsParamsRegionMena,
	})
	if err != nil {
		var apierr *y2.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
