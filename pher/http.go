package pher

import (
	"net/url"
	"os"
)

func Get(key string) string {
	url := &url.URL{RawQuery: os.Getenv("QUERY_STRING")}
	return url.Query().Get(key)
}
