package usecases

import (
	"github.com/co0p/go-tls-watch/pkg/domain"
)

type Fetcher interface {
	Fetch(string) (domain.Certificate, error)
}

type ValidateUsecase struct {
	Client Fetcher
}

func (v *ValidateUsecase) Validate(website string) (domain.Certificate, error) {
	return v.Client.Fetch(website)
}
