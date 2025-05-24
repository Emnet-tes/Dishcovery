package utils

import (
	"net/http"
	"os"

	"github.com/machinebox/graphql"
)

type headersTransport struct {
	headers http.Header
	base    http.RoundTripper
}

func (t *headersTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, values := range t.headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	return t.base.RoundTrip(req)
}

var client *graphql.Client

func Client() *graphql.Client {
	if client == nil {
		hasuraURL := os.Getenv("HASURA_GRAPHQL_URL")
		if hasuraURL == "" {
			hasuraURL = "http://localhost:8080/v1/graphql" // Ensure URL is correct
		}

		httpClient := &http.Client{
			Transport: &headersTransport{
				headers: http.Header{
					"x-hasura-role": []string{"admin"},
				},
				base: http.DefaultTransport,
			},
		}

		client = graphql.NewClient(hasuraURL, graphql.WithHTTPClient(httpClient))
	}

	return client
}