package api

import (
	"net/http"
	"net/url"
)

type IAPI interface {
	Get(url, body string, params url.Values, header http.Header) ([]byte, int, error)
}
