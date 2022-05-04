package net

import (
	"fmt"
	"net/http"
)

// Client is a structure that wraps a Http Client
type Client struct {
	http.Client
}

// Do is a wrapper for Do function in http.Client
func (rc Client) Do(req *http.Request) (*http.Response, error) {
	// Implement helper method to keep headers
	rc.Client.CheckRedirect = rc.checkRedirect

	// Continue http.Client Do method
	return rc.Client.Do(req)
}

// checkRedirect is a helper method that keeps Request Headers for every redirect
func (Client) checkRedirect(req *http.Request, via []*http.Request) error {
	// Set the limit of redirects to 15 before you stop saving the sensitive headers
	if len(via) > 15 {
		return fmt.Errorf("%d consecutive redirects", len(via))
	}
	if len(via) == 0 {
		return nil
	}

	// Go through each key and value and add it back to the request headers for the next redirect
	for key, val := range via[0].Header {
		req.Header[key] = val
	}
	return nil
}
