package mock

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

// Test used to test the Mock Client
func TestClientMock_Do(t *testing.T) {

	// Create a list of urls to test and append them to a slice
	var urlStrings = []string{
		"https://mcmp-api-stage.cisco.com/v1/tenants/faketenant/servicegroups",
	}
	var urls []*url.URL
	for _, u := range urlStrings {
		u, _ := url.Parse(u)
		urls = append(urls, u)
	}

	// Create your table test
	var tests = []struct {
		name     string
		method   string
		url      string
		response http.Response
		request  http.Request
	}{
		{
			name:   "MCMP Add SG Test 1",
			method: "POST",
			url:    urls[0].String(),
			response: http.Response{
				Status:     "201 Resource Created",
				StatusCode: 201,
				Header: http.Header{
					"test": {"This is a test response"},
				},
			},
			request: http.Request{
				Method: "POST",
				URL:    urls[0],
				Header: http.Header{
					"X-Tenant-Id":  {"tenant id"},
					"Content-Type": {"application/vnd.cia.v1+json"},
					"Accept":       {"application/vnd.cia.v1+json"},
				},
				Body: nil,
			},
		},
	}

	// Go through each of the tests in the table
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Create the mock Client
			client := Client{}

			// Add the responses to the client
			client.SetResponses(tt.method, tt.url, tt.response)

			// Run the Do method on the client and make sure you get the correct response back
			got, _ := client.Do(&tt.request)
			if !reflect.DeepEqual(got, &tt.response) {
				t.Errorf("wanted %v, got %v", tt.response, got)
			}
		})
	}
}
