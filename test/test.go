package main

import (
	"io"
	"log"

	http "github.com/bogdanfinn/fhttp"
	tlsclient "github.com/bogdanfinn/tls-client"

	"github.com/xiaozhou26/tlsclient/tls_client"
)

func main() {
	// Create a cookie jar for managing cookies
	jar := tlsclient.NewCookieJar()

	// Set up the HTTP client options, including using the Chrome127 profile
	options := []tlsclient.HttpClientOption{
		tlsclient.WithTimeoutSeconds(30),
		tlsclient.WithClientProfile(tls_client.Chrome127()), // Use your custom Chrome127 profile
		tlsclient.WithNotFollowRedirects(),
		tlsclient.WithCookieJar(jar),
	}

	// Create the HTTP client
	client, err := tlsclient.NewHttpClient(tlsclient.NewNoopLogger(), options...)
	if err != nil {
		log.Fatalf("Failed to create HTTP client: %v", err)
	}

	// Create a new HTTP GET request
	req, err := http.NewRequest(http.MethodGet, "https://tls.peet.ws/api/all", nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Set request headers
	req.Header = http.Header{
		"accept":          {"*/*"},
		"accept-language": {"de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7"},
		"user-agent":      {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"},
		http.HeaderOrderKey: {
			"accept",
			"accept-language",
			"user-agent",
		},
	}

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	// Log the status code
	log.Printf("Status code: %d", resp.StatusCode)

	// Read and log the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	log.Printf("Response body: %s", body)
}
