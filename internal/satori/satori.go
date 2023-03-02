package satori

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type SatoriClient struct {
	client     *http.Client
	url        *url.URL
	urlString  string
	apiKey     string
	signingKey string
}

func NewSatoriClient(satoriUrl string, apiKey string, signingKey string) *SatoriClient {
	parsedUrl, _ := url.Parse(satoriUrl)
	return &SatoriClient{
		urlString:  satoriUrl,
		client:     &http.Client{Timeout: 2 * time.Second},
		url:        parsedUrl,
		apiKey:     strings.TrimSpace(apiKey),
		signingKey: strings.TrimSpace(signingKey),
	}
}

func (s *SatoriClient) validateConfig() error {
	var errUrl, errApiKey, errSigningKey error
	if s.url == nil {
		_, err := url.Parse(s.urlString)
		errUrl = fmt.Errorf("Invalid Satori URL: %s", err.Error())
	}
	if s.apiKey == "" {
		errApiKey = errors.New("Satori API Key not set.")
	}
	if s.signingKey == "" {
		errSigningKey = errors.New("Satori Signing Key is not set.")
	}

	if err := errors.Join(errUrl, errApiKey, errSigningKey); err != nil {
		return err
	}

	return nil
}

func (s *SatoriClient) generateToken() string {

}

func (s *SatoriClient) Foo() {

}
