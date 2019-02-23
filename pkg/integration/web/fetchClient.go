package web

import "crypto/tls"

type FetchClient struct {
}

func (c *FetchClient) Fetch(website string) (tls.Certificate, error) {
	return tls.Certificate{}, nil
}
