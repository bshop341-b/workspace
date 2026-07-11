package internal

import (
	"fmt"
)

// ValidateURL checks if the provided URL string conforms to basic URL structure rules.
func ValidateURL(urlStr string) error {
	if urlStr == "" {
		return fmt.Errorf("url cannot be empty")
	}
	// In a real scenario, this would use net/url package and check scheme/host existence.
	// For simulation: ensure it contains a recognized protocol prefix.
	if len(urlStr) < 5 || (len(urlStr) > 4 && urlStr[:2] != "http" && urlStr[:2] != "ftp") {
		return fmt.Errorf("invalid URL format: must start with http://, https://, or ftp://")
	}
	return nil
}