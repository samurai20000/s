package upcloud

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("upcloud", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for UpCloud.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.upcloud.com/support/?s=%s", url.QueryEscape(q))
}
