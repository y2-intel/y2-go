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

func TestProfileNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Profiles.New(context.TODO(), y2.ProfileNewParams{
		Frequency:         y2.ProfileNewParamsFrequencyDaily,
		Name:              "Cybersecurity Weekly",
		ScheduleTimeOfDay: "09:00",
		Topic:             "Cybersecurity threats, vulnerabilities, and defense strategies",
		BlufStructure:     y2.String("blufStructure"),
		CustomPrompt:      y2.String("customPrompt"),
		IsCommunity:       y2.Bool(true),
		RecursionConfig: y2.ProfileNewParamsRecursionConfig{
			Enabled:  true,
			MaxDepth: 1,
			Strategy: "breadth-first",
		},
		ScheduleDayOfMonth: y2.String("1"),
		ScheduleDayOfWeek:  y2.String("monday"),
		Tags:               []string{"string"},
	})
	if err != nil {
		var apierr *y2.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProfileUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Profiles.Update(
		context.TODO(),
		"k57abc123def456",
		y2.ProfileUpdateParams{
			AudioConfig: y2.ProfileUpdateParamsAudioConfig{
				Enabled: y2.Bool(true),
				Speed:   y2.Float(0),
				VoiceID: y2.String("voiceId"),
			},
			BlufStructure:      y2.String("blufStructure"),
			BrandingTemplateID: y2.String("brandingTemplateId"),
			BudgetConfig: y2.ProfileUpdateParamsBudgetConfig{
				AlertThreshold:   y2.Float(0),
				MaxCostPerReport: y2.Float(0),
			},
			CustomPrompt: y2.String("customPrompt"),
			Frequency:    y2.ProfileUpdateParamsFrequencyDaily,
			FreshnessConfig: y2.ProfileUpdateParamsFreshnessConfig{
				Enabled:             y2.Bool(true),
				MaxAgeMs:            y2.Int(0),
				PreferRecentSources: y2.Bool(true),
				RecencyWeight:       y2.Float(0),
				ValidateLinks:       y2.Bool(true),
			},
			IsCommunity: y2.Bool(true),
			ModelConfig: y2.ProfileUpdateParamsModelConfig{
				MaxOutputTokens: y2.Int(0),
				ModelID:         y2.String("modelId"),
				Temperature:     y2.Float(0),
			},
			Name: y2.String("name"),
			RecursionConfig: y2.ProfileUpdateParamsRecursionConfig{
				Enabled:  true,
				MaxDepth: 1,
				Strategy: "breadth-first",
			},
			ScheduleDayOfMonth: y2.String("scheduleDayOfMonth"),
			ScheduleDayOfWeek:  y2.String("scheduleDayOfWeek"),
			ScheduleTimeOfDay:  y2.String("73:16"),
			SearchConfig: y2.ProfileUpdateParamsSearchConfig{
				ExcludeDomains: []string{"string"},
				IncludeDomains: []string{"string"},
				MaxResults:     y2.Int(0),
				SearchDepth:    "basic",
				TimeRange:      y2.String("timeRange"),
				Topic:          y2.String("topic"),
			},
			Status: y2.ProfileUpdateParamsStatusActive,
			Tags:   []string{"string"},
			Topic:  y2.String("topic"),
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

func TestProfileList(t *testing.T) {
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
	_, err := client.Profiles.List(context.TODO())
	if err != nil {
		var apierr *y2.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProfileDelete(t *testing.T) {
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
	_, err := client.Profiles.Delete(context.TODO(), "k57abc123def456")
	if err != nil {
		var apierr *y2.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProfilePartialUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Profiles.PartialUpdate(
		context.TODO(),
		"k57abc123def456",
		y2.ProfilePartialUpdateParams{
			AudioConfig: y2.ProfilePartialUpdateParamsAudioConfig{
				Enabled: y2.Bool(true),
				Speed:   y2.Float(0),
				VoiceID: y2.String("voiceId"),
			},
			BlufStructure:      y2.String("blufStructure"),
			BrandingTemplateID: y2.String("brandingTemplateId"),
			BudgetConfig: y2.ProfilePartialUpdateParamsBudgetConfig{
				AlertThreshold:   y2.Float(0),
				MaxCostPerReport: y2.Float(0),
			},
			CustomPrompt: y2.String("customPrompt"),
			Frequency:    y2.ProfilePartialUpdateParamsFrequencyDaily,
			FreshnessConfig: y2.ProfilePartialUpdateParamsFreshnessConfig{
				Enabled:             y2.Bool(true),
				MaxAgeMs:            y2.Int(0),
				PreferRecentSources: y2.Bool(true),
				RecencyWeight:       y2.Float(0),
				ValidateLinks:       y2.Bool(true),
			},
			IsCommunity: y2.Bool(true),
			ModelConfig: y2.ProfilePartialUpdateParamsModelConfig{
				MaxOutputTokens: y2.Int(0),
				ModelID:         y2.String("modelId"),
				Temperature:     y2.Float(0),
			},
			Name: y2.String("name"),
			RecursionConfig: y2.ProfilePartialUpdateParamsRecursionConfig{
				Enabled:  true,
				MaxDepth: 1,
				Strategy: "breadth-first",
			},
			ScheduleDayOfMonth: y2.String("scheduleDayOfMonth"),
			ScheduleDayOfWeek:  y2.String("scheduleDayOfWeek"),
			ScheduleTimeOfDay:  y2.String("73:16"),
			SearchConfig: y2.ProfilePartialUpdateParamsSearchConfig{
				ExcludeDomains: []string{"string"},
				IncludeDomains: []string{"string"},
				MaxResults:     y2.Int(0),
				SearchDepth:    "basic",
				TimeRange:      y2.String("timeRange"),
				Topic:          y2.String("topic"),
			},
			Status: y2.ProfilePartialUpdateParamsStatusActive,
			Tags:   []string{"string"},
			Topic:  y2.String("topic"),
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
