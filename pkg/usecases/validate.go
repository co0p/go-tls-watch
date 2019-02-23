package usecases

import (
	"crypto/tls"

	"github.com/co0p/go-tls-watch/pkg/domain"
)

type Fetcher interface {
	Fetch(string) (tls.Certificate, error)
}

type ValidateUsecase struct {
	Client Fetcher
}

func (v *ValidateUsecase) Validate(website string) (domain.CertificateInfo, error) {
	return domain.CertificateInfo{}, nil
}
