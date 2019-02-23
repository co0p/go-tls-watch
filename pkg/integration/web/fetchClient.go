package web

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/co0p/go-tls-watch/pkg/domain"
)

type FetchClient struct {
	client *http.Client
}

func NewFetchClient() FetchClient {

	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	return FetchClient{
		client: client,
	}
}

func (c *FetchClient) Fetch(website string) (domain.Certificate, error) {
	res, err := c.client.Get(website)

	if err != nil {
		log.Println("failed fetching certificate: ", err)
		return domain.Certificate{}, err
	}

	if res.TLS == nil {
		return domain.Certificate{}, errors.New("no certificate found")
	}

	cert := res.TLS.PeerCertificates[0]
	issuer := cert.Issuer.String()
	notBefore := cert.NotBefore
	notAfter := cert.NotAfter

	return domain.Certificate{
		Origin:    website,
		Issuer:    issuer,
		NotBefore: notBefore,
		NotAfter:  notAfter,
	}, nil
}
