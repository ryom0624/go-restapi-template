package lib

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	baseUri = "https://www.google.com/recaptcha/api/siteverify"

	defaultMinimumVerificationScore = 0.7
)

type Recaptcha interface {
	Verify(token string) error
}

type RecaptchaCli struct {
	Secret string `json:"secret"`
}

type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

func NewRecaptchaCli(secret string) (*RecaptchaCli, error) {
	//if secret == "" {
	//	return nil, errors.New("recaptcha secret is empty")
	//}

	return &RecaptchaCli{Secret: secret}, nil
}

func (c *RecaptchaCli) Verify(token string) error {
	resp, err := http.PostForm(baseUri,
		url.Values{"secret": []string{c.Secret}, "response": {token}})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var r RecaptchaResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return err
	}

	if !r.Success {
		return errors.New("recaptcha verification failed")
	}
	if r.Score < defaultMinimumVerificationScore {
		return errors.New("recaptcha verification failed")
	}

	return nil
}

type NoopRecaptchaCli struct{}

func NewNoopRecaptchaCli(secret string) (*NoopRecaptchaCli, error) {
	return &NoopRecaptchaCli{}, nil
}

func (c *NoopRecaptchaCli) Verify(token string) error {
	return nil
}
