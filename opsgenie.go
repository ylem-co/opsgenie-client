package opsgenieclient

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type Opsgenie struct {
	Client *resty.Client
	ctx context.Context
}

func (o *Opsgenie) CreateAlertWithContext(ctx context.Context, request CreateAlertRequest) error {
	const MessageMaxLength = 130
	if len(request.Message) > MessageMaxLength {
		return &ValidationError{
			Err: fmt.Errorf("length of a message may not exceed %d characters", MessageMaxLength),
		}
	}

	const DescriptionMaxLength = 15000
	if len(request.Description) > DescriptionMaxLength {
		return &ValidationError{
			Err: fmt.Errorf("length of a description may not exceed %d characters", DescriptionMaxLength),
		}
	}

	if _, ok := map[string]bool{
		"P1": true,
		"P2": true,
		"P3": true,
		"P4": true,
		"P5": true,
	}[request.Priority]; !ok {
		return &ValidationError{
			Err: fmt.Errorf("a priority %s is not supported", request.Priority),
		}
	}

	resp, err := o.Client.R().
	SetBody(request).
	SetContext(ctx).
	Post("/v2/alerts")

	if err != nil {
		return fmt.Errorf("opsgenie: %s", err.Error())
	}

	if resp.StatusCode() != http.StatusAccepted {
		return fmt.Errorf(
			"opsgenie: Expected: %d http response got %s. Body: %s",
			http.StatusAccepted,
			resp.Status(),
			string(resp.Body()),
		)
	}

	return nil
}

func (o *Opsgenie) CreateAlert(request CreateAlertRequest) error {
	return o.CreateAlertWithContext(o.ctx, request)
}

func CreateInstance(ctx context.Context, apiKey string) (*Opsgenie, error) {
	return &Opsgenie{
		Client: resty.New().
			SetBaseURL("https://api.eu.opsgenie.com").
			SetAuthScheme("GeniKey").
			SetAuthToken(apiKey),
		ctx: ctx,
	}, nil
}
