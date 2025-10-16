package github

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/v63/github"
)

// Client is a wrapper around GitHub client that supports GitHub App authentication for multiple installations.
type Client struct {
	*github.Client

	transport *ghinstallation.AppsTransport
}

type CustomRoundTripper struct {
	Proxied http.RoundTripper
}

func (crt *CustomRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header["Accept"] = []string{"application/vnd.github.machine-man-preview+json"}
	return crt.Proxied.RoundTrip(req)
}

type BaseURLRoundTripper struct {
	Proxied http.RoundTripper
	BaseURL string
}

func (b *BaseURLRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if b.BaseURL != "" {
		parsedBaseURL, err := url.Parse(b.BaseURL)
		if err != nil {
			return nil, fmt.Errorf("parsing base URL in BaseURLRoundTripper: %w", err)
		}
		req.URL.Scheme = parsedBaseURL.Scheme
		req.URL.Host = parsedBaseURL.Host
	}
	return b.Proxied.RoundTrip(req)
}

// NewClient creates a new Client.
func NewClient(appID int64, appPrivateKey string, baseURL string, insecureSkipVerify bool) (*Client, error) {
	// Ensure baseURL for ghinstallation does NOT have a trailing slash
	ghinstallationBaseURL := strings.TrimSuffix(baseURL, "/")

	// Ensure baseURL for go-github client DOES have a trailing slash
	goGithubBaseURL := baseURL
	if !strings.HasSuffix(goGithubBaseURL, "/") {
		goGithubBaseURL += "/"
	}

	// Create a custom http.Transport to handle insecureSkipVerify
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureSkipVerify},
	}

	// Wrap the transport with BaseURLRoundTripper to set the BaseURL for ghinstallation's internal client
	baseURLRoundTripper := &BaseURLRoundTripper{Proxied: transport, BaseURL: ghinstallationBaseURL}

	appTransport, err := ghinstallation.NewAppsTransport(baseURLRoundTripper, appID, []byte(appPrivateKey))
	if err != nil {
		return nil, err
	}

	parsedGoGithubBaseURL, err := url.Parse(goGithubBaseURL)
	if err != nil {
		return nil, fmt.Errorf("parsing go-github base URL: %w", err)
	}

	// appTransport.BaseURL is already set by BaseURLRoundTripper, no need to set it again here
	// appTransport.BaseURL = ghinstallationBaseURL

	// Wrap the appTransport with CustomRoundTripper to add the Accept header
	customTransport := &CustomRoundTripper{Proxied: appTransport}

	client := github.NewClient(&http.Client{Transport: customTransport})

	client.BaseURL = parsedGoGithubBaseURL

	return &Client{
		Client:    client,
		transport: appTransport,
	},
	nil
}

// Installation returns a new GitHub client for the given installation ID.
func (c *Client) Installation(installationID int64) *github.Client {
	installationTransport := ghinstallation.NewFromAppsTransport(c.transport, installationID)
	// Ensure BaseURL for installationTransport does NOT have a trailing slash
	installationTransport.BaseURL = strings.TrimSuffix(c.BaseURL.String(), "/")
	customInstallationTransport := &CustomRoundTripper{Proxied: installationTransport}
	installationClient := github.NewClient(&http.Client{Transport: customInstallationTransport})
	installationClient.BaseURL = c.BaseURL
	return installationClient
}
