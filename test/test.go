package main

import (
	"io"
	"log"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/xiaozhou26/tlsclient/tlsclient"
)

func main() {
	// Create a cookie jar for managing cookies
	jar := tls_client.NewCookieJar()

	// Set up the HTTP client options, including using the Chrome127 profile
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(30),
		tls_client.WithClientProfile(tlsclient.Safari()), // Use your custom Chrome127 profile
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar),
	}

	// Create the HTTP client
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		log.Fatalf("Failed to create HTTP client: %v", err)
	}
	client.SetProxy("http://127.0.0.1:2080")
	// Create a new HTTP GET request
	req, err := http.NewRequest(http.MethodGet, "https://chatgpt.com/", nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Set request headers
	req.Header = http.Header{
		"accept": {
			"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		},
		"accept-language": {
			"zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7",
		},
		"cache-control": {
			"max-age=0",
		},
		"priority": {
			"u=0, i",
		},
		"sec-ch-ua": {
			"\"Not)A;Brand\";v=\"99\", \"Google Chrome\";v=\"127\", \"Chromium\";v=\"127\"",
		},
		"sec-ch-ua-arch": {
			"\"x86\"",
		},
		"sec-ch-ua-bitness": {
			"\"64\"",
		},
		"sec-ch-ua-full-version": {
			"\"127.0.6533.120\"",
		},
		"sec-ch-ua-full-version-list": {
			"\"Not)A;Brand\";v=\"99.0.0.0\", \"Google Chrome\";v=\"127.0.6533.120\", \"Chromium\";v=\"127.0.6533.120\"",
		},
		"sec-ch-ua-mobile": {
			"?0",
		},
		"sec-ch-ua-model": {
			"\"\"",
		},
		"sec-ch-ua-platform": {
			"\"Windows\"",
		},
		"sec-ch-ua-platform-version": {
			"\"15.0.0\"",
		},
		"sec-fetch-dest": {
			"document",
		},
		"sec-fetch-mode": {
			"navigate",
		},
		"sec-fetch-site": {
			"same-origin",
		},
		"sec-fetch-user": {
			"?1",
		},
		"upgrade-insecure-requests": {
			"1",
		},
		"user-agent": {
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36",
		},
		http.HeaderOrderKey: {
			"accept",
			"accept-language",
			"cache-control",
			"priority",
			"sec-ch-ua",
			"sec-ch-ua-arch",
			"sec-ch-ua-bitness",
			"sec-ch-ua-full-version",
			"sec-ch-ua-full-version-list",
			"sec-ch-ua-mobile",
			"sec-ch-ua-model",
			"sec-ch-ua-platform",
			"sec-ch-ua-platform-version",
			"sec-fetch-dest",
			"sec-fetch-mode",
			"sec-fetch-site",
			"sec-fetch-user",
			"upgrade-insecure-requests",
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

	readBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(string(readBytes))

}
