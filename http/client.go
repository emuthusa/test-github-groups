package http

import (
	"io"
	"net/http"
)

// A Client interface that wraps the net.http client's Do method
type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

// Response is a wrapper for the http.Response
type Response http.Response

// NewRequest is used to wrap the http.NewRequest method
func NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}

// NewResponse copies a http.Response into our Response object
func NewResponse(res *http.Response) *Response {
	return &Response{
		Status:           res.Status,
		StatusCode:       res.StatusCode,
		Proto:            res.Proto,
		ProtoMajor:       res.ProtoMajor,
		ProtoMinor:       res.ProtoMinor,
		Header:           res.Header,
		Body:             res.Body,
		ContentLength:    res.ContentLength,
		TransferEncoding: res.TransferEncoding,
		Close:            res.Close,
		Uncompressed:     res.Uncompressed,
		Trailer:          res.Trailer,
		Request:          res.Request,
		TLS:              res.TLS,
	}
}
